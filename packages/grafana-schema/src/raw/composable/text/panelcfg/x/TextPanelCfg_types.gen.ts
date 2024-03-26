// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     public/app/plugins/gen.go
// Using jennies:
//     TSTypesJenny
//     LatestMajorsOrXJenny
//     PluginEachMajorJenny
//
// Run 'make gen-cue' from repository root to regenerate.

export const pluginVersion = "10.1.10";

export enum TextMode {
  Code = 'code',
  HTML = 'html',
  Markdown = 'markdown',
}

export enum CodeLanguage {
  Go = 'go',
  Html = 'html',
  Json = 'json',
  Markdown = 'markdown',
  Plaintext = 'plaintext',
  Sql = 'sql',
  Typescript = 'typescript',
  Xml = 'xml',
  Yaml = 'yaml',
}

export const defaultCodeLanguage: CodeLanguage = CodeLanguage.Plaintext;

export interface CodeOptions {
  /**
   * The language passed to monaco code editor
   */
  language: CodeLanguage;
  showLineNumbers: boolean;
  showMiniMap: boolean;
}

export const defaultCodeOptions: Partial<CodeOptions> = {
  language: CodeLanguage.Plaintext,
  showLineNumbers: false,
  showMiniMap: false,
};

export interface Options {
  code?: CodeOptions;
  content: string;
  mode: TextMode;
}

export const defaultOptions: Partial<Options> = {
  content: `# Title

For markdown syntax help: [commonmark.org/help](https://commonmark.org/help/)`,
  mode: TextMode.Markdown,
};
