export default defineNuxtRouteMiddleware((to, from) => {
  // この middleware が設定されている場合は要ログイン
  const { currentUser } = useAuth();
  if (!currentUser.value && to.path !== "/login") {
    const path = "/login";
    return { path };
  }
});
