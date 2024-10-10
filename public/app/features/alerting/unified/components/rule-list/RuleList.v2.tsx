import { css } from '@emotion/css';
import { PropsWithChildren, ReactNode, useMemo } from 'react';
import Skeleton from 'react-loading-skeleton';

import { GrafanaTheme2 } from '@grafana/data';
import { Icon, LinkButton, Pagination, Stack, Text, useStyles2, withErrorBoundary } from '@grafana/ui';
import { Rule, RuleGroupIdentifier, RuleIdentifier } from 'app/types/unified-alerting';
import { RulesSourceApplication } from 'app/types/unified-alerting-dto';

import { alertRuleApi } from '../../api/alertRuleApi';
import { featureDiscoveryApi } from '../../api/featureDiscoveryApi';
import { getAllRulesSources, isGrafanaRulesSource } from '../../utils/datasource';
import { equal, fromRule, fromRulerRule, hashRule, stringifyIdentifier } from '../../utils/rule-id';
import { getRulePluginOrigin, isAlertingRule, isRecordingRule } from '../../utils/rules';
import { createRelativeUrl } from '../../utils/url';
import { AlertingPageWrapper } from '../AlertingPageWrapper';
import { Spacer } from '../Spacer';
import { WithReturnButton } from '../WithReturnButton';
import RulesFilter from '../rules/Filter/RulesFilter';
import { ActionsLoader, RuleActionsButtons } from '../rules/RuleActionsButtons';

import { AlertRuleListItem, RecordingRuleListItem, UnknownRuleListItem } from './AlertRuleListItem';
import { DataSourceIcon } from './Namespace';
import { ListGroup } from './components/ListGroup';
import { ListSection } from './components/ListSection';
import { LoadingIndicator } from './components/RuleGroup';

const noop = () => {};
const { usePrometheusRuleNamespacesQuery, useGetRuleGroupForNamespaceQuery } = alertRuleApi;

const RuleList = withErrorBoundary(
  () => {
    const ruleSources = getAllRulesSources();

    return (
      // We don't want to show the Loading... indicator for the whole page.
      // We show separate indicators for Grafana-managed and Cloud rules
      <AlertingPageWrapper navId="alert-list" isLoading={false} actions={null}>
        <RulesFilter onClear={() => {}} />
        <Stack direction="column" gap={1}>
          {ruleSources.map((ruleSource) => {
            if (isGrafanaRulesSource(ruleSource)) {
              return <GrafanaDataSourceLoader key={ruleSource} />;
            } else {
              return <DataSourceLoader key={ruleSource.uid} uid={ruleSource.uid} name={ruleSource.name} />;
            }
          })}
        </Stack>
      </AlertingPageWrapper>
    );
  },
  { style: 'page' }
);

const { useDiscoverDsFeaturesQuery } = featureDiscoveryApi;

interface DataSourceLoaderProps {
  name: string;
  uid: string;
}

const GrafanaDataSourceLoader = () => {
  return <DataSourceSection name="Grafana" application="grafana" isLoading={true}></DataSourceSection>;
};

const DataSourceLoader = ({ uid, name }: DataSourceLoaderProps) => {
  const { data: dataSourceInfo, isLoading } = useDiscoverDsFeaturesQuery({ uid });
  let application: RulesSourceApplication | undefined;

  if (dataSourceInfo?.dataSourceSettings.type === 'loki') {
    application = 'loki';
  } else {
    application = dataSourceInfo?.features.application;
  }

  if (isLoading) {
    return <DataSourceSection loader={<Skeleton width={250} height={16} />}></DataSourceSection>;
  }

  // 2. grab prometheus rule groups with max_groups if supported
  if (dataSourceInfo) {
    const rulerEnabled = Boolean(dataSourceInfo.rulerConfig);

    return (
      <DataSourceSection name={name} application={application} uid={uid}>
        <PaginatedRuleGroupLoader
          ruleSourceName={dataSourceInfo?.dataSourceSettings.name}
          rulerEnabled={rulerEnabled}
        />
      </DataSourceSection>
    );
  }

  return null;
};

interface PaginatedRuleGroupLoaderProps {
  ruleSourceName: string;
  rulerEnabled?: boolean;
}

function PaginatedRuleGroupLoader({ ruleSourceName, rulerEnabled = false }: PaginatedRuleGroupLoaderProps) {
  const { data: ruleNamespaces = [], isLoading } = usePrometheusRuleNamespacesQuery({
    ruleSourceName,
    maxGroups: 25,
    limitAlerts: 0,
    excludeAlerts: true,
  });

  return (
    <Stack direction="column" gap={1}>
      {ruleNamespaces.map((namespace) => (
        <ListSection
          key={namespace.name}
          title={
            <Stack direction="row" gap={1} alignItems="center">
              <Icon name="folder" /> {namespace.name}
            </Stack>
          }
        >
          {namespace.groups.map((group) => (
            <ListGroup key={group.name} name={group.name} isOpen={false}>
              {group.rules.map((rule) => {
                const groupIdentifier: RuleGroupIdentifier = {
                  dataSourceName: ruleSourceName,
                  groupName: group.name,
                  namespaceName: namespace.name,
                };

                return (
                  <AlertRuleLoader
                    key={hashRule(rule)}
                    rule={rule}
                    groupIdentifier={groupIdentifier}
                    rulerEnabled={rulerEnabled}
                  />
                );
              })}
            </ListGroup>
          ))}
        </ListSection>
      ))}
      {!isLoading && <Pagination currentPage={1} numberOfPages={0} onNavigate={noop} />}
    </Stack>
  );
}

interface AlertRuleLoaderProps {
  rule: Rule;
  groupIdentifier: RuleGroupIdentifier;
  rulerEnabled?: boolean;
}

function AlertRuleLoader({ rule, groupIdentifier, rulerEnabled = false }: AlertRuleLoaderProps) {
  const { dataSourceName, namespaceName, groupName } = groupIdentifier;

  const ruleIdentifier = fromRule(dataSourceName, namespaceName, groupName, rule);
  const href = createViewLinkFromIdentifier(ruleIdentifier);
  const originMeta = getRulePluginOrigin(rule);

  // @TODO work with context API to propagate rulerConfig and such
  const { data: dataSourceInfo } = useDiscoverDsFeaturesQuery({ rulesSourceName: dataSourceName });

  // @TODO refactor this to use a separate hook (useRuleWithLocation() and useCombinedRule() seems to introduce infinite loading / recursion)
  const {
    isLoading,
    data: rulerRuleGroup,
    error,
  } = useGetRuleGroupForNamespaceQuery(
    {
      namespace: namespaceName,
      group: groupName,
      rulerConfig: dataSourceInfo?.rulerConfig!,
    },
    { skip: !dataSourceInfo?.rulerConfig }
  );

  const rulerRule = useMemo(() => {
    if (!rulerRuleGroup) {
      return;
    }

    return rulerRuleGroup.rules.find((rule) =>
      equal(fromRulerRule(dataSourceName, namespaceName, groupName, rule), ruleIdentifier)
    );
  }, [dataSourceName, groupName, namespaceName, ruleIdentifier, rulerRuleGroup]);

  // 1. get the rule from the ruler API with "ruleWithLocation"
  // 1.1 skip this if this datasource does not have a ruler
  //
  // 2.1 render action buttons
  // 2.2 render provisioning badge and contact point metadata, etc.

  const actions = useMemo(() => {
    if (!rulerEnabled) {
      return null;
    }

    if (isLoading) {
      return <ActionsLoader />;
    }

    if (rulerRule) {
      return <RuleActionsButtons rule={rulerRule} promRule={rule} groupIdentifier={groupIdentifier} compact />;
    }

    return null;
  }, [groupIdentifier, isLoading, rule, rulerEnabled, rulerRule]);

  if (isAlertingRule(rule)) {
    return (
      <AlertRuleListItem
        name={rule.name}
        href={href}
        summary={rule.annotations?.summary}
        state={rule.state}
        health={rule.health}
        error={rule.lastError}
        labels={rule.labels}
        isProvisioned={undefined}
        instancesCount={undefined}
        actions={actions}
        origin={originMeta}
      />
    );
  }

  if (isRecordingRule(rule)) {
    return (
      <RecordingRuleListItem
        name={rule.name}
        href={href}
        health={rule.health}
        error={rule.lastError}
        labels={rule.labels}
        isProvisioned={undefined}
        actions={null}
        origin={originMeta}
      />
    );
  }

  return <UnknownRuleListItem rule={rule} groupIdentifier={groupIdentifier} />;
}

function createViewLinkFromIdentifier(identifier: RuleIdentifier, returnTo?: string) {
  const paramId = encodeURIComponent(stringifyIdentifier(identifier));
  const paramSource = encodeURIComponent(identifier.ruleSourceName);

  return createRelativeUrl(`/alerting/${paramSource}/${paramId}/view`, returnTo ? { returnTo } : {});
}

interface DataSourceSectionProps extends PropsWithChildren {
  uid?: string;
  name?: string;
  loader?: ReactNode;
  application?: RulesSourceApplication;
  isLoading?: boolean;
}

const DataSourceSection = ({ uid, name, application, children, loader, isLoading = false }: DataSourceSectionProps) => {
  const styles = useStyles2(getStyles);

  return (
    <Stack direction="column" gap={1}>
      <Stack direction="column" gap={0}>
        {isLoading && <LoadingIndicator />}
        <div className={styles.dataSourceSectionTitle}>
          {loader ?? (
            <Stack alignItems="center">
              {application && <DataSourceIcon application={application} />}
              {name && (
                <Text variant="body" weight="bold">
                  {name}
                </Text>
              )}
              <Spacer />
              {uid && (
                <WithReturnButton
                  title="alert rules"
                  component={
                    <LinkButton variant="secondary" size="sm" href={`/connections/datasources/edit/${uid}`}>
                      Configure
                    </LinkButton>
                  }
                />
              )}
            </Stack>
          )}
        </div>
      </Stack>
      <div className={styles.itemsWrapper}>{children}</div>
    </Stack>
  );
};

const getStyles = (theme: GrafanaTheme2) => ({
  itemsWrapper: css({
    position: 'relative',
    marginLeft: theme.spacing(1.5),

    '&:before': {
      content: "''",
      position: 'absolute',
      height: '100%',

      marginLeft: `-${theme.spacing(1.5)}`,
      borderLeft: `solid 1px ${theme.colors.border.weak}`,
    },
  }),
  dataSourceSectionTitle: css({
    background: theme.colors.background.secondary,
    padding: `${theme.spacing(1)} ${theme.spacing(1.5)}`,

    border: `solid 1px ${theme.colors.border.weak}`,
    borderRadius: theme.shape.radius.default,
  }),
});

export default RuleList;
