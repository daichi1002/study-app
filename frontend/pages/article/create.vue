<script setup lang="ts">
import { Article } from "~/types/article";
import { changePage } from "~/util/router";

const { article } = useArticle();

const saveArticle = async () => {
  // TODO:もっとリッチにしたい
  if (!article.value.title) return alert("タイトルを入力してください");
  if (!article.value.title) return alert("本文を入力してください");

  const { error } = await useFetch<Article>(
    "http://localhost:8080/article/create",
    {
      method: "POST",
      body: article,
    }
  );

  if (error.value) {
    throw createError({
      statusCode: 404,
      message: "failed to create",
    });
  }

  changePage("/article/list");
};
</script>

<template>
  <div class="container mx-auto">
    <Form />
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
