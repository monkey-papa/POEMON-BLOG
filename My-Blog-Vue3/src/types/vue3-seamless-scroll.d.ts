/**
 * Type declarations for vue3-seamless-scroll
 * Overrides the library's incorrect type definitions
 * The library requires ComputedGetter for props, but normal values work fine in Vue 3
 */

// Override the vue3-seamless-scroll module types
declare module "vue3-seamless-scroll" {
  import type { DefineComponent } from "vue";

  // Props accept both regular values and computed/ref values
  type PropValue<T> = T | (() => T);

  interface SeamlessScrollProps {
    modelValue?: PropValue<boolean>;
    list: PropValue<unknown[]>;
    step?: PropValue<number>;
    hover?: PropValue<boolean>;
    direction?: PropValue<"up" | "down" | "left" | "right">;
    singleWaitTime?: PropValue<number>;
    ease?: PropValue<"ease-in" | "linear" | "ease" | "ease-out" | "ease-in-out" | string>;
    wheel?: PropValue<boolean>;
    visibleCount?: PropValue<number>;
    singleLine?: PropValue<boolean>;
  }

  export const Vue3SeamlessScroll: DefineComponent<SeamlessScrollProps>;

  const install: (app: import("vue").App, options?: { name?: string }) => void;
  export default install;
}
