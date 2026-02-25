module.exports = {
  root: true,
  env: {
    node: true,
    browser: true,
    es2021: true,
  },
  extends: [
    "plugin:vue/vue3-essential",
    "plugin:@typescript-eslint/recommended",
    "plugin:prettier/recommended",
  ],
  parser: "vue-eslint-parser",
  parserOptions: {
    parser: "@typescript-eslint/parser",
    ecmaVersion: 2021,
    sourceType: "module",
    extraFileExtensions: [".vue"],
  },
  globals: {
    // 全局库和工具
    $: "readonly", // jQuery
    NProgress: "readonly", // 进度条库
    tocbot: "readonly", // 目录生成库
    hljs: "readonly", // 代码高亮库
    APlayer: "readonly", // 音乐播放器
    anime: "readonly", // 动画库
    SakuraList: "readonly", // 樱花效果
    Hls: "readonly", // HLS 视频播放
    ClipboardJS: "readonly", // 剪贴板库
    iziToast: "readonly", // 提示消息库
    currentMusic: "readonly", // 当前音乐对象
    define: "readonly", // AMD 模块定义（用于第三方库）
    // TypeScript 全局类型
    JQueryStatic: "readonly",
    JQuery: "readonly",
    NodeListOf: "readonly",
    NodeJS: "readonly",
    EventListener: "readonly",
  },
  plugins: ["@typescript-eslint"],
  rules: {
    // TypeScript 未使用变量检测
    "@typescript-eslint/no-unused-vars": [
      "warn",
      {
        argsIgnorePattern: "^_",
        varsIgnorePattern: "^_",
        ignoreRestSiblings: true,
      },
    ],
    "@typescript-eslint/no-empty-object-type": "off",
    "no-unused-vars": "off", // 使用 TypeScript 版本替代
    "@typescript-eslint/no-non-null-assertion": "off",
    "@typescript-eslint/no-var-requires": "off",
    "no-eval": "off",
    "no-var": "warn",
    "no-debugger": "warn", // 禁用debugger
    "no-duplicate-case": "warn", // 禁止出现重复的 case 标签
    "no-empty": "warn", // 禁止出现空语句块
    "no-extra-parens": "off", // 禁止不必要的括号
    "no-func-assign": "warn", // 禁止对 function 声明重新赋值
    "no-unreachable": "warn", // 禁止在 return、throw、continue 和 break 语句之后出现不可达代码
    "default-case": "warn", // 要求 switch 语句中有 default 分支
    "dot-notation": "warn", // 强制尽可能地使用点号
    eqeqeq: "off", // 要求使用 === 和 !==
    "no-else-return": "off", // 禁止 if 语句中 return 语句之后有 else 块
    "no-sequences": "off",
    "prefer-promise-reject-errors": "off",
    "no-empty-function": "off", // 禁止出现空函数
    "array-callback-return": "off",
    "no-unused-expressions": "off",
    "no-new": "off",
    "prefer-regex-literals": "off",
    "no-use-before-define": "off",
    "one-var": "off",
    "no-lone-blocks": "warn", // 禁用不必要的嵌套块
    "no-multi-spaces": "warn", // 禁止使用多个空格
    "no-redeclare": "warn", // 禁止多次声明同一变量
    "no-return-assign": "warn", // 禁止在 return 语句中使用赋值语句
    "no-return-await": "warn", // 禁用不必要的 return await
    "no-self-assign": "warn", // 禁止自我赋值
    "no-self-compare": "warn", // 禁止自身比较
    "no-useless-catch": "warn", // 禁止不必要的 catch 子句
    "no-useless-return": "warn", // 禁止多余的 return 语句
    "no-shadow": "off", // 禁止变量声明与外层作用域的变量同名
    "no-delete-var": "off", // 允许delete变量
    "array-bracket-spacing": "warn", // 强制数组方括号中使用一致的空格
    "brace-style": "warn", // 强制在代码块中使用一致的大括号风格
    camelcase: "warn", // 强制使用骆驼拼写法命名约定
    // indent: ["error", 2], // tab缩进
    "jsx-quotes": "warn", // 强制在 JSX 属性中一致地使用双引号或单引号
    "max-depth": "warn", // 强制可嵌套的块的最大深度4
    "max-lines": ["warn", { max: 3000 }], // 强制最大行数 300
    "max-statements": ["warn", 100], // 强制函数块最多允许的的语句数量20
    "max-nested-callbacks": ["warn", 10], // 强制回调函数最大嵌套深度
    "max-depth": ["warn", 10],
    "max-statements-per-line": ["warn", { max: 1 }], // 强制每一行中所允许的最大语句数量
    "newline-per-chained-call": ["warn", { ignoreChainWithDepth: 3 }], // 要求方法链中每个调用都有一个换行符
    "no-lonely-if": "off", // 禁止 if 作为唯一的语句出现在 else 语句中
    "no-mixed-spaces-and-tabs": "warn", // 禁止空格和 tab 的混合缩进
    "no-multiple-empty-lines": "warn", // 禁止出现多行空行
    // 禁止出现;
    // 结尾分号
    semi: "off",
    // 双引号
    quotes: "off",
    // 强制在块之前使用一致的空格
    "space-before-blocks": "warn",
    // 强制在 function的左括号之前使用一致的空格
    // 'space-before-function-paren': ['warn', 'never'],
    // 强制在圆括号内使用一致的空格
    "space-in-parens": "warn",
    // 要求操作符周围有空格
    "space-infix-ops": "warn",
    // 强制在一元操作符前后使用一致的空格
    "space-unary-ops": "warn",
    // 强制在注释中 // 或 /* 使用一致的空格
    "spaced-comment": "off",
    // 强制在 switch 的冒号左右有空格
    "switch-colon-spacing": "warn",
    // 强制箭头函数的箭头前后使用一致的空格
    "arrow-spacing": "warn",
    "prefer-const": "warn",
    "prefer-rest-params": "warn",
    "no-useless-escape": "off",
    "no-irregular-whitespace": "warn",
    "no-prototype-builtins": "warn",
    "no-fallthrough": "warn",
    "no-extra-boolean-cast": "warn",
    "no-case-declarations": "off",
    "no-async-promise-executor": "off",
    "comma-dangle": 0,
    "no-undef": "error",
    // vue rule
    "no-loss-of-precision": "off",
    "no-nonoctal-decimal-escape": "off",
    "no-unsafe-optional-chaining": "off",
    "no-useless-backreference": "off",
    "no-dupe-keys": "off",
    "no-unused-vars": "off",
    "newline-per-chained-call": "off",
    "vue/valid-attribute-name": "off",
    "vue/valid-model-definition": "off",
    "vue/no-v-html": "off",
    "vue/attributes-order": "off",
    "vue/require-default-prop": "off",
    "vue/require-prop-type-constructor": "off",
    "vue/no-lone-template": "off",
    "vue/no-mutating-props": "off",
    "vue/no-useless-template-attributes": "off",
    "vue/no-dupe-keys": "off",
    "vue/no-v-model-argument": "off",
    "vue/no-unused-components": "warn",
    "vue/no-side-effects-in-computed-properties": "warn",
    "vue/html-self-closing": [
      "error",
      {
        html: {
          void: "any",
          normal: "any",
          component: "always",
        },
        svg: "always",
        math: "always",
      },
    ],
    // 多字组件名称
    "vue/multi-word-component-names": "off",
    "vue/max-attributes-per-line": [
      "error",
      {
        singleline: {
          max: 10,
        },
        multiline: {
          max: 10,
        },
      },
    ],
    "@typescript-eslint/no-empty-function": "off",
    "@typescript-eslint/no-this-alias": "off",
    "@typescript-eslint/no-explicit-any": "warn",
    camelcase: "off",
    "@typescript-eslint/no-non-null-assertion": "off",
    "node/no-callback-literal": "off",
  },
  // 忽略第三方库文件
  ignorePatterns: [
    "src/utils/APlayer.min.js",
    "src/utils/sakura.ts",
    "src/utils/sakura.js",
    "src/utils/mousedown.js",
  ],
};
