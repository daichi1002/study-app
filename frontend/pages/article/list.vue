<script setup lang="ts">
import { Article } from "~/types/article";

// const config = await useRuntimeConfig();
const { data: articles } = await useFetch<Article[]>(
  `http://localhost:8080/article/list`
);
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
          <tr v-for="article in articles">
            <td class="border px-4 py-2">
              <NuxtLink :to="`/article/${article.articleId}`">{{
                article.title
              }}</NuxtLink>
            </td>
            <td class="border px-4 py-2">{{ article.tag }}</td>
            <td class="border px-4 py-2">{{ article.userName }}</td>
            <td class="border px-4 py-2">{{ article.updatedDate }}</td>
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