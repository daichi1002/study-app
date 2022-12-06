<script setup lang="ts">
import { Article } from "~/types/article";

const reqArticle = reactive<Article>({
  articleId: null,
  title: "",
  content: "",
  tags: [],
  updatedAt: null,
});

const saveArticle = async () => {
  const { data, error } = await useFetch<Article>(
    "http://localhost:8080/article/create",
    {
      method: "POST",
      body: reqArticle,
    }
  );

  if (error.value) {
    throw createError({
      statusCode: 404,
      message: "failed to create",
    });
  }
};
</script>

<template>
  <div class="container mx-auto">
    <Form :value="reqArticle" />
    <div class="mt-4 grid grid-cols-9 gap-9">
      <button
        class="bg-teal-300 hover:bg-teal-500 py-2 px-4 rounded-lg"
        @click="saveArticle()"
      >
        保存
      </button>
    </div>
  </div>
</template>
