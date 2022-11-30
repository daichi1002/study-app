<script setup lang="ts">
import { Tag } from "~/types/tag";
import { RequestArticle } from "~/types/article";

const { data: tags, error } = await useFetch<Tag[]>(
  `http://localhost:8080/create`
);

if (!tags.value || error.value) {
  throw createError({
    statusCode: 404,
    message: "Tags not found",
  });
}
const reqArticle = reactive<RequestArticle>({
  title: "",
  content: "",
  tagId: [],
  userId: "ABCDEFGHIJK", // user管理機能未作成なので仮の値を設定
});

const saveArticle = async () => {
  const { data, error } = await useFetch<RequestArticle>(
    "http://localhost:8080/create",
    {
      method: "POST",
      body: reqArticle,
    }
  );
};
</script>

<template>
  <div class="container mx-auto">
    <div class="mt-5">
      <div>タイトル</div>
      <input
        class="appearance-none border-2 border-gray-300 rounded w-2/4 py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
        id="inline-full-name"
        type="text"
        v-model="reqArticle.title"
      />
    </div>
    <div class="mt-4">
      <div>タグ</div>
      <div class="grid grid-cols-5 gap-5">
        <div v-for="tag in tags">
          <input
            id="default-checkbox"
            type="checkbox"
            :value="tag.id"
            v-model="reqArticle.tagId"
            class="w-4 h-4 text-blue-600 bg-gray-100 rounded border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
          />
          <label for="default-checkbox" class="ml-2 text-sm font-medium">{{
            tag.name
          }}</label>
        </div>
      </div>
    </div>
    <div class="mt-4">
      <div>本文</div>
      <!-- <textarea
        class="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full h-72 py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
        id="inline-full-name"
        type="text"
        value=""
      /> -->
      <Markdown v-model="reqArticle.content" />
    </div>
    <div class="mt-4 grid grid-cols-9 gap-9">
      <button
        class="bg-teal-300 hover:bg-teal-500 py-2 px-4 rounded-lg"
        @click="saveArticle()"
      >
        保存
      </button>
      <button class="bg-red-400 hover:bg-red-500 py-2 px-4 rounded-lg">
        削除
      </button>
    </div>
  </div>
</template>
