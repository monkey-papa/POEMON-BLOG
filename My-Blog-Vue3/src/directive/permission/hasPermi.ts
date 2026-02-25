import { useStore } from "@/stores";
import type { Directive, DirectiveBinding } from "vue";

interface PermiDirectiveBinding extends DirectiveBinding<string[]> {
  arg?: "disabled";
}

const hasPermi: Directive<HTMLElement, string[]> = {
  mounted(el: HTMLElement, binding: PermiDirectiveBinding) {
    const { value, arg } = binding;
    const all_permission = "*:*:*";
    const store = useStore();
    const permissions = store.permissions;
    if (value && value instanceof Array && value.length > 0) {
      const permissionFlag = value;
      const hasPermissions = permissions.some((permission: string) => {
        return all_permission === permission || permissionFlag.includes(permission);
      });
      if (!hasPermissions) {
        if (arg === "disabled") {
          el.classList.add("is-disabled");
          el.setAttribute("disabled", "true");
        } else {
          if (el.parentNode) {
            el.parentNode.removeChild(el);
          }
        }
      }
    } else {
      throw new Error("请设置操作权限标签值");
    }
  },
};

export default hasPermi;
