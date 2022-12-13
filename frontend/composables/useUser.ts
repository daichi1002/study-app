import type { Ref } from "vue";
import { User } from "~/types/user";

export const useUser = () => {
  const users = useState<User[]>("users", () => []);

  const setUsers = (resUsers: User[] | null) => {
    if (resUsers != null) {
      users.value = resUsers;
    }
  };

  return {
    users: readonly(users),
    setUsers,
  };
};
