<script setup lang="ts">
import { Article } from "~/types/article";
import { changePage } from "~/util/router";

const { article, resetArticle } = useArticle();
resetArticle();

const saveArticle = async () => {
  // TODO:もっとリッチにしたい
  if (!article.value.title) return alert("タイトルを入力してください");
  if (!article.value.content) return alert("本文を入力してください");

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
  <div class="container max-w-8xl px-4 mx-auto sm:px-8">
    <Form />
    <div class="mt-4 grid grid-cols-12 gap-9">
      <button
        class="flex-shrink-0 px-4 py-2 text-base text-white bg-emerald-300 rounded-lg shadow-md hover:bg-emerald-500 focus:outline-none"
        @click="saveArticle()"
      >
        保存
      </button>
    </div>
  </div>
</template>
