import store from "@/store";
export default {
  inserted(el, binding) {
    const { value, arg } = binding;
    const all_permission = "*:*:*";
    const permissions = store.getters && store.getters.permissions;
    if (value && value instanceof Array && value.length > 0) {
      const permissionFlag = value;
      const hasPermissions = permissions.some((permission) => {
        return (
          all_permission === permission || permissionFlag.includes(permission)
        );
      });
      if (!hasPermissions) {
        if (arg === "disabled") {
          el.classList.add("is-disabled");
          el.setAttribute("disabled", true);
        } else {
          el.parentNode && el.parentNode.removeChild(el);
        }
      }
    } else {
      throw new Error("请设置操作权限标签值");
    }
  },
};
