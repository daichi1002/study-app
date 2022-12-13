<script setup lang="ts">
import { Article } from "~/types/article";
import { User } from "~/types/user";
import { formatDate, setTagName, setUserName } from "~/util/format";

const { data: articles } = await useFetch<Article[]>(
  `http://localhost:8080/article/list`
);

const { data: users } = await useFetch<User[]>("http://localhost:8080/users");

const { setUsers } = useUser();
setUsers(unref(users));
</script>
<template>
  <div class="container mx-auto">
    <div class="flex justify-center mt-10">
      <input
        class="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-2/4 py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
        id="inline-full-name"
        type="text"
        value=""
      />
      <button>
        <img src="../../assets/images/search.png" />
      </button>
    </div>
    <div class="flex justify-center mt-10">
      <table class="table-auto">
        <thead>
          <tr>
            <th class="px-4 py-2">タイトル</th>
            <th class="px-4 py-2">タグ</th>
            <th class="px-4 py-2">投稿者</th>
            <th class="px-4 py-2">更新日</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="article in articles" class="hover:bg-slate-100">
            <td class="border px-4 py-2">
              <NuxtLink :to="`/article/${article.articleId}`">{{
                article.title
              }}</NuxtLink>
            </td>
            <td class="border px-4 py-2">{{ setTagName(article.tags) }}</td>
            <!--  ユーザー機能ができ次第表示（現在は仮の値） -->
            <td class="border px-4 py-2">{{ setUserName(article.userId) }}</td>
            <td class="border px-4 py-2">
              {{ formatDate(article.updatedAt) }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <NuxtLink to="/article/create">
      <button
        class="fixed bottom-8 right-8 rounded-full w-24 py-6 bg-teal-300 hover:bg-teal-500 text-white text-center font-semibold text-5xl"
      >
        +
      </button>
    </NuxtLink>
  </div>
</template>
