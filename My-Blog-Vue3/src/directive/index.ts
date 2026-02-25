import hasPermi from "./permission/hasPermi";
import type { App } from "vue";

export default {
  install(app: App): void {
    app.directive("hasPermi", hasPermi);
  },
};
