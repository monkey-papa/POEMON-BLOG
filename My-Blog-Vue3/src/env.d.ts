/// <reference types="vite/client" />

declare module "*.vue" {
  import type { DefineComponent } from "vue";
  const component: DefineComponent<object, object, unknown>;
  export default component;
}

// ==================== 第三方库模块声明 ====================

declare module "jquery" {
  const $: JQueryStatic;
  export default $;
}

declare module "animejs" {
  const anime: AnimeStatic;
  export default anime;
}

declare module "wow.js" {
  export default class WOW {
    constructor(options?: {
      boxClass?: string;
      animateClass?: string;
      offset?: number;
      mobile?: boolean;
      live?: boolean;
      callback?: (box: HTMLElement) => void;
      scrollContainer?: string | null;
      resetAnimation?: boolean;
    });
    init(): void;
    sync(): void;
  }
}

declare module "nprogress" {
  const nprogress: NProgressStatic;
  export default nprogress;
}

declare module "aplayer" {
  export default class APlayer {
    constructor(options: {
      container: HTMLElement;
      fixed?: boolean;
      mini?: boolean;
      autoplay?: boolean;
      theme?: string;
      loop?: "all" | "one" | "none";
      order?: "list" | "random";
      preload?: "none" | "metadata" | "auto";
      volume?: number;
      audio: APlayerAudio | APlayerAudio[];
      mutex?: boolean;
      lrcType?: number;
      listFolded?: boolean;
      listMaxHeight?: number | string;
      storageName?: string;
    });
    play(): void;
    pause(): void;
    toggle(): void;
    seek(time: number): void;
    volume(percentage?: number): number;
    on(event: string, handler: () => void): void;
    destroy(): void;
    list: {
      add(audios: APlayerAudio | APlayerAudio[]): void;
      remove(index: number): void;
      switch(index: number): void;
      clear(): void;
    };
  }

  interface APlayerAudio {
    name: string;
    artist: string;
    url: string;
    cover: string;
    lrc?: string;
    theme?: string;
    type?: string;
  }
}

declare module "colorthief" {
  export default class ColorThief {
    getColor(img: HTMLImageElement | null, quality?: number): [number, number, number] | null;
    getPalette(
      img: HTMLImageElement | null,
      colorCount?: number,
      quality?: number
    ): [number, number, number][] | null;
  }
}

declare module "tocbot" {
  interface TocbotOptions {
    tocSelector: string;
    contentSelector: string;
    headingSelector: string;
    hasInnerContainers?: boolean;
    linkClass?: string;
    extraLinkClasses?: string;
    activeLinkClass?: string;
    listClass?: string;
    extraListClasses?: string;
    isCollapsedClass?: string;
    collapsibleClass?: string;
    listItemClass?: string;
    collapseDepth?: number;
    scrollSmooth?: boolean;
    scrollSmoothDuration?: number;
    scrollSmoothOffset?: number;
    headingsOffset?: number;
    throttleTimeout?: number;
    positionFixedSelector?: string;
    positionFixedClass?: string;
    fixedSidebarOffset?: string | number;
    includeHtml?: boolean;
    onClick?: (e: Event) => void;
    orderedList?: boolean;
    scrollContainer?: string | null;
    skipRendering?: boolean;
    headingLabelCallback?: (headingLabel: string) => string;
    ignoreHiddenElements?: boolean;
    headingObjectCallback?: (obj: object, element: HTMLElement) => object;
    basePath?: string;
    disableTocScrollSync?: boolean;
  }

  const tocbot: {
    init(options: TocbotOptions): void;
    refresh(options?: TocbotOptions): void;
    destroy(): void;
  };
  export default tocbot;
}

declare module "clipboard" {
  export default class ClipboardJS {
    constructor(
      selector: string | Element | NodeListOf<Element>,
      options?: {
        target?: (trigger: Element) => Element;
        text?: (trigger: Element) => string;
        container?: Element;
        action?: (trigger: Element) => "copy" | "cut";
      }
    );
    on(
      type: "success" | "error",
      handler: (e: {
        action: "copy" | "cut";
        text: string;
        trigger: Element;
        clearSelection(): void;
      }) => void
    ): this;
    destroy(): void;
    static copy(text: string): string;
    static cut(target: Element): string;
    static isSupported(): boolean;
  }
}

declare module "highlightjs-line-numbers.js";

declare module "markdown-it" {
  interface MarkdownItOptions {
    html?: boolean;
    xhtmlOut?: boolean;
    breaks?: boolean;
    langPrefix?: string;
    linkify?: boolean;
    typographer?: boolean;
    quotes?: string;
    highlight?: (str: string, lang: string) => string;
  }

  class MarkdownIt {
    constructor(presetName?: "commonmark" | "zero" | "default", options?: MarkdownItOptions);
    constructor(options?: MarkdownItOptions);
    render(md: string, env?: object): string;
    renderInline(md: string, env?: object): string;
    parse(src: string, env?: object): unknown[];
    parseInline(src: string, env?: object): unknown[];
    use(plugin: unknown, ...params: unknown[]): this;
    disable(rules: string | string[], ignoreInvalid?: boolean): this;
    enable(rules: string | string[], ignoreInvalid?: boolean): this;
    set(options: MarkdownItOptions): this;
  }

  export default MarkdownIt;
}

// ==================== 全局类型定义 ====================

// Anime.js 类型
interface AnimeInstance {
  play(): void;
  pause(): void;
  restart(): void;
  reverse(): void;
  seek(time: number): void;
  finished: Promise<void>;
}

interface AnimeTimelineInstance extends AnimeInstance {
  add(params: AnimeParams, offset?: string | number): AnimeTimelineInstance;
}

interface AnimeParams {
  targets?: unknown;
  duration?: number;
  delay?: number | ((el: Element, index: number, total: number) => number);
  easing?: string;
  elasticity?: number;
  round?: number | boolean;
  loop?: number | boolean;
  direction?: "normal" | "reverse" | "alternate";
  autoplay?: boolean;
  update?: (anim: AnimeCallbackParam) => void;
  begin?: (anim: AnimeCallbackParam) => void;
  complete?: (anim: AnimeCallbackParam) => void;
  loopBegin?: (anim: AnimeCallbackParam) => void;
  loopComplete?: (anim: AnimeCallbackParam) => void;
  changeBegin?: (anim: AnimeCallbackParam) => void;
  changeComplete?: (anim: AnimeCallbackParam) => void;
  [key: string]: unknown;
}

interface AnimeCallbackParam {
  animatables: Array<{ target: unknown }>;
  progress: number;
  currentTime: number;
  remaining: number;
  completed: boolean;
  paused: boolean;
}

interface AnimeStatic {
  (params: AnimeParams): AnimeInstance;
  timeline(params?: AnimeParams): AnimeTimelineInstance;
  random(min: number, max: number): number;
  stagger(value: number | string, options?: object): unknown;
  set(targets: unknown, props: object): void;
  remove(targets: unknown): void;
  get(targets: unknown, prop: string, unit?: string): unknown;
  path(path: string | Element): unknown;
  setDashoffset(el: Element): number;
  bezier(x1: number, y1: number, x2: number, y2: number): (t: number) => number;
  running: AnimeInstance[];
  easings: Record<string, (t: number) => number>;
  version: string;
}

// HLjs 语言信息类型
interface HljsLanguage {
  name: string;
  aliases?: string[];
  case_insensitive?: boolean;
  keywords?: object;
}

// HLjs 类型
interface HljsStatic {
  highlightAll(): void;
  highlightElement(element: Element): void;
  highlightBlock(block: Element): void;
  lineNumbersBlock(block: Element): void;
  initLineNumbersOnLoad?(): void;
  configure(options: object): void;
  getLanguage(name: string): HljsLanguage | undefined;
  listLanguages(): string[];
  registerLanguage(name: string, language: object): void;
  highlight(
    code: string,
    options: { language: string; ignoreIllegals?: boolean }
  ): { value: string; language: string };
  highlightAuto(code: string, languageSubset?: string[]): { value: string; language?: string };
}

// NProgress 类型
interface NProgressStatic {
  configure(options: {
    minimum?: number;
    template?: string;
    easing?: string;
    speed?: number;
    trickle?: boolean;
    trickleSpeed?: number;
    showSpinner?: boolean;
    parent?: string;
    positionUsing?: string;
  }): NProgressStatic;
  set(number: number): NProgressStatic;
  start(): NProgressStatic;
  done(force?: boolean): NProgressStatic;
  inc(amount?: number): NProgressStatic;
  remove(): void;
  isStarted(): boolean;
  isRendered(): boolean;
  status: number | null;
  version: string;
}

// Tocbot 类型
interface TocbotStatic {
  init(options: {
    tocSelector: string;
    contentSelector: string;
    headingSelector: string;
    hasInnerContainers?: boolean;
    linkClass?: string;
    activeLinkClass?: string;
    listClass?: string;
    listItemClass?: string;
    collapseDepth?: number;
    scrollSmooth?: boolean;
    scrollSmoothDuration?: number;
    scrollSmoothOffset?: number;
    headingsOffset?: number;
    throttleTimeout?: number;
    positionFixedSelector?: string;
    positionFixedClass?: string;
    fixedSidebarOffset?: string | number;
    includeHtml?: boolean;
    onClick?: (e: Event) => void;
    orderedList?: boolean;
    scrollContainer?: string | null;
  }): void;
  refresh(options?: object): void;
  destroy(): void;
}

// ClipboardJS 类型
interface ClipboardJSStatic {
  new (
    selector: string | Element | NodeListOf<Element>,
    options?: {
      target?: (trigger: Element) => Element;
      text?: (trigger: Element) => string;
      container?: Element;
      action?: (trigger: Element) => "copy" | "cut";
    }
  ): ClipboardJSInstance;
  copy(text: string): string;
  cut(target: Element): string;
  isSupported(): boolean;
}

interface ClipboardJSInstance {
  on(
    type: "success" | "error",
    handler: (e: {
      action: "copy" | "cut";
      text: string;
      trigger: Element;
      clearSelection(): void;
    }) => void
  ): this;
  destroy(): void;
}

// WOW 类型
interface WOWStatic {
  new (options?: {
    boxClass?: string;
    animateClass?: string;
    offset?: number;
    mobile?: boolean;
    live?: boolean;
    callback?: (box: HTMLElement) => void;
    scrollContainer?: string | null;
    resetAnimation?: boolean;
  }): WOWInstance;
}

interface WOWInstance {
  init(): void;
  sync(): void;
}

// ColorThief 类型
interface ColorThiefStatic {
  new (): ColorThiefInstance;
}

interface ColorThiefInstance {
  getColor(img: HTMLImageElement | null, quality?: number): [number, number, number] | null;
  getPalette(
    img: HTMLImageElement | null,
    colorCount?: number,
    quality?: number
  ): [number, number, number][] | null;
}

// APlayer 类型
interface APlayerStatic {
  new (options: {
    container: HTMLElement | null;
    fixed?: boolean | string;
    mini?: boolean;
    autoplay?: boolean;
    theme?: string;
    loop?: "all" | "one" | "none";
    order?: "list" | "random";
    preload?: "none" | "metadata" | "auto";
    volume?: number;
    audio: APlayerAudio | APlayerAudio[];
    mutex?: boolean;
    lrcType?: number;
    listFolded?: boolean;
    listMaxHeight?: number | string;
    storageName?: string;
  }): APlayerInstance;
}

interface APlayerAudio {
  name: string;
  artist: string;
  url: string;
  cover: string;
  lrc?: string;
  theme?: string;
  type?: string;
}

interface APlayerInstance {
  play(): void;
  pause(): void;
  toggle(): void;
  seek(time: number): void;
  volume(percentage?: number): number;
  on(event: string, handler: () => void): void;
  skipBack(): void;
  skipForward(): void;
  destroy(): void;
  list: {
    current: {
      title: string;
      author: string;
    };
    add(audios: APlayerAudio | APlayerAudio[]): void;
    remove(index: number): void;
    switch(index: number): void;
    clear(): void;
    show(): void;
    hide(): void;
  };
}

// iziToast 类型
interface IziToastOptions {
  timeout?: number;
  icon?: string;
  displayMode?: "replace" | "once" | "replaceAll";
  message?: string;
  title?: string;
  position?: string;
  color?: string;
  backgroundColor?: string;
  progressBar?: boolean;
  progressBarColor?: string;
  image?: string;
  imageWidth?: number;
  maxWidth?: number;
  zindex?: number;
  layout?: number;
  balloon?: boolean;
  close?: boolean;
  closeOnEscape?: boolean;
  closeOnClick?: boolean;
  rtl?: boolean;
  pauseOnHover?: boolean;
  resetOnHover?: boolean;
  drag?: boolean;
  transitionIn?: string;
  transitionOut?: string;
  transitionInMobile?: string;
  transitionOutMobile?: string;
  buttons?: Array<{
    text: string;
    color?: string;
    close?: boolean;
    callback?: () => void;
  }>;
  inputs?: Array<{
    name: string;
    placeholder?: string;
    value?: string;
  }>;
  onOpen?: () => void;
  onClose?: () => void;
}

interface IziToastStatic {
  show(options: IziToastOptions): void;
  success(options: IziToastOptions): void;
  error(options: IziToastOptions): void;
  info(options: IziToastOptions): void;
  warning(options: IziToastOptions): void;
  question(options: IziToastOptions): void;
  progress(options: IziToastOptions & { progress?: number }): void;
  hide(options?: { transitionOut?: string }, toast?: HTMLElement): void;
  destroy(): void;
  settings(options: IziToastOptions): void;
}

// ==================== 全局 Window 扩展 ====================

declare global {
  interface Window {
    $: JQueryStatic;
    jQuery: JQueryStatic;
    anime: AnimeStatic;
    hljs: HljsStatic;
    WOW: WOWStatic;
    ClipboardJS: ClipboardJSStatic;
    NProgress: NProgressStatic;
    echarts: typeof import("echarts/core");
    APlayer: APlayerStatic;
    ColorThief: ColorThiefStatic;
    tocbot: TocbotStatic;
    iziToast: IziToastStatic;
    hideInitialLoading?: () => void;
    startSakura?: () => void;
  }

  // 全局变量
  const $: JQueryStatic;
  const jQuery: JQueryStatic;
  const NProgress: NProgressStatic;
  const tocbot: TocbotStatic;
  const hljs: HljsStatic;
  const APlayer: APlayerStatic;
  const anime: AnimeStatic;
}

// ==================== Vue 全局属性扩展 ====================

import type { ElMessageBox } from "element-plus";

declare module "@vue/runtime-core" {
  interface ComponentCustomProperties {
    $common: typeof import("./utils/common").default;
    $constant: typeof import("./utils/constant").default;
    $confirm: typeof ElMessageBox.confirm;
  }
}

// ==================== 路由元信息扩展 ====================

import "vue-router";

declare module "vue-router" {
  interface RouteMeta {
    requiresAuth?: boolean;
    title?: string;
  }
}

export {};
