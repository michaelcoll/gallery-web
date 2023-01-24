/* eslint-env node */
require('@rushstack/eslint-patch/modern-module-resolution')

module.exports = {
  root: true,
  'plugins': [
    'simple-import-sort'
  ],
  'extends': [
    'plugin:vue/vue3-recommended',
    'eslint:recommended',
    '@vue/eslint-config-typescript',
    '@vue/eslint-config-prettier'
  ],
  'rules': {
    'simple-import-sort/imports': 'error',
    'simple-import-sort/exports': 'error'
  },
  parserOptions: {
    ecmaVersion: 'latest'
  }
}
