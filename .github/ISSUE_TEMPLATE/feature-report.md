name: Feature Report
about: Suggest an idea for this project
title: ''
labels: feature
assignees: ''
  - 
body:
  - type: textarea
    id: about-feature
    attributes:
      label: 新規機能の概要
      description: 新規機能の概要を書いてください.
      placeholder: 〇〇の開発をする.
    validations:
      required: true
  - type: textarea
    id: version
    attributes:
      label: バージョン
      description: リリース予定バージョン
      placeholder: version X.X.X
    validations:
      required: false
