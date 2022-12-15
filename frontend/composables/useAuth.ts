import type { Ref } from "vue";

const login = (currentUser: Ref<boolean>) => async () => {
  currentUser.value = true;
  // middleware でログインページにリダイレクトした場合は redirectFrom に元のページが入っている
  const to = useRoute().redirectedFrom?.path || "/";
  useRouter().push(to);
};

const logout = (currentUser: Ref<boolean>) => async () => {
  currentUser.value = false;
};

export const useAuth = () => {
  const currentUser = useState("currentUser", () => false);
  return {
    currentUser,
    login: login(currentUser),
    logout: logout(currentUser),
  };
};
