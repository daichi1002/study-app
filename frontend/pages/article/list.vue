<script setup lang="ts">
import { Article } from "~/types/article";
import { User } from "~/types/user";
import { formatDate, setTagName, setUserName } from "~/util/format";

const { data: articles } = await useFetch<Article[]>(
  `http://localhost:8080/article/list`
);

const { data: users } = await useFetch<User[]>("http://localhost:8080/users");
</script>
<template>
  <div class="container max-w-8xl px-4 mx-auto sm:px-8">
    <div class="py-8">
      <div class="flex flex-row justify-between w-full mb-1 sm:mb-0">
        <h2 class="text-2xl leading-tight">投稿記事</h2>
        <div class="text-end">
          <form
            class="flex flex-col justify-center w-3/4 max-w-sm space-y-3 md:flex-row md:w-full md:space-x-3 md:space-y-0"
          >
            <div class="relative">
              <input
                type="text"
                class="rounded-lg border-transparent flex-1 appearance-none border border-gray-200 w-full py-2 px-4 bg-white text-gray-700 placeholder-gray-400 shadow-sm text-base focus:outline-none focus:ring-2 focus:ring-emerald-300 focus:border-transparent"
                placeholder="タイトル"
              />
            </div>
            <button
              class="flex-shrink-0 px-4 py-2 text-base text-white bg-emerald-300 rounded-lg shadow-md hover:bg-emerald-500 focus:outline-none"
              type="submit"
            >
              検索
            </button>
          </form>
        </div>
      </div>
      <div class="px-4 py-4 -mx-4 overflow-x-auto sm:-mx-8 sm:px-8">
        <div
          class="inline-block min-w-full overflow-hidden rounded-lg shadow-md"
        >
          <table class="min-w-full leading-normal">
            <thead>
              <tr>
                <th
                  scope="col"
                  class="px-5 py-3 text-sm font-normal text-left text-gray-800 uppercase bg-white border-b border-gray-300"
                >
                  タイトル
                </th>
                <th
                  scope="col"
                  class="px-5 py-3 text-sm font-normal text-left text-gray-800 uppercase bg-white border-b border-gray-300"
                >
                  タグ
                </th>
                <th
                  scope="col"
                  class="px-5 py-3 text-sm font-normal text-left text-gray-800 uppercase bg-white border-b border-gray-300"
                >
                  更新者
                </th>
                <th
                  scope="col"
                  class="px-5 py-3 text-sm font-normal text-left text-gray-800 uppercase bg-white border-b border-gray-300"
                >
                  更新日
                </th>
                <!-- <th
                  scope="col"
                  class="px-5 py-3 text-sm font-normal text-left text-gray-800 uppercase bg-white border-b border-gray-200"
                ></th> -->
              </tr>
            </thead>
            <tbody>
              <tr v-for="article in articles" class="hover:bg-slate-100">
                <td class="px-5 py-5 text-sm bg-white border-b border-gray-300">
                  <NuxtLink :to="`/article/${article.articleId}`">{{
                    article.title
                  }}</NuxtLink>
                </td>
                <td class="px-5 py-5 text-sm bg-white border-b border-gray-300">
                  {{ setTagName(article.tags) }}
                </td>
                <!--  ユーザー機能ができ次第表示（現在は仮の値） -->
                <td class="px-5 py-5 text-sm bg-white border-b border-gray-300">
                  {{ setUserName(article.userId, users) }}
                </td>
                <td class="px-5 py-5 text-sm bg-white border-b border-gray-300">
                  {{ formatDate(article.updatedAt) }}
                </td>
              </tr>
            </tbody>
          </table>
          <div
            class="flex flex-col items-center px-5 py-5 bg-white xs:flex-row xs:justify-between"
          >
            <div class="flex items-center">
              <button
                type="button"
                class="w-full p-4 text-base text-gray-600 bg-white border rounded-l-xl hover:bg-gray-100"
              >
                <svg
                  width="9"
                  fill="currentColor"
                  height="8"
                  class=""
                  viewBox="0 0 1792 1792"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M1427 301l-531 531 531 531q19 19 19 45t-19 45l-166 166q-19 19-45 19t-45-19l-742-742q-19-19-19-45t19-45l742-742q19-19 45-19t45 19l166 166q19 19 19 45t-19 45z"
                  ></path>
                </svg>
              </button>
              <button
                type="button"
                class="w-full px-4 py-2 text-base text-indigo-500 bg-white border-t border-b hover:bg-gray-100"
              >
                1
              </button>
              <button
                type="button"
                class="w-full px-4 py-2 text-base text-gray-600 bg-white border hover:bg-gray-100"
              >
                2
              </button>
              <button
                type="button"
                class="w-full px-4 py-2 text-base text-gray-600 bg-white border-t border-b hover:bg-gray-100"
              >
                3
              </button>
              <button
                type="button"
                class="w-full px-4 py-2 text-base text-gray-600 bg-white border hover:bg-gray-100"
              >
                4
              </button>
              <button
                type="button"
                class="w-full p-4 text-base text-gray-600 bg-white border-t border-b border-r rounded-r-xl hover:bg-gray-100"
              >
                <svg
                  width="9"
                  fill="currentColor"
                  height="8"
                  class=""
                  viewBox="0 0 1792 1792"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M1363 877l-742 742q-19 19-45 19t-45-19l-166-166q-19-19-19-45t19-45l531-531-531-531q-19-19-19-45t19-45l166-166q19-19 45-19t45 19l742 742q19 19 19 45t-19 45z"
                  ></path>
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <NuxtLink to="/article/create">
    <img
      class="fixed bottom-8 right-8 w-20"
      src="../../assets/images/plus.png"
    />
  </NuxtLink>
</template>
