<script setup lang="ts">
import { Article } from "~/types/article";
import { changePage } from "~/util/router";

const router = useRoute();
const { id } = router.params;

const { data } = await useFetch<Article>(`http://localhost:8080/article/${id}`);
const { article, setArticle } = useArticle();
const props = setArticle(data.value);

// TODO:複数回PUTリクエストが送られているので、見直し要
const updateArticle = async () => {
  // TODO:もっとリッチにしたい
  if (!article.value.title) return alert("タイトルを入力してください");
  if (!article.value.content) return alert("本文を入力してください");

  const { error } = await useFetch<Article>(
    `http://localhost:8080/article/edit/${id}`,
    {
      method: "PUT",
      body: article,
    }
  );

  if (error.value) {
    throw createError({
      statusCode: 404,
      message: "failed to create",
    });
  }

  changePage(`/article/${id}`);
};

const deleteArticle = async () => {
  const { error } = await useFetch<Article>(
    `http://localhost:8080/article/delete/${id}`,
    {
      method: "DELETE",
    }
  );

  if (error.value) {
    throw createError({
      statusCode: 404,
      message: "failed to delete",
    });
  }

  changePage("/article/list");
};
</script>

<template>
  <div class="container max-w-8xl px-4 mx-auto sm:px-8">
    <Form :article="props" />
    <div class="mt-4 grid grid-cols-12 gap-9">
      <button
        class="flex-shrink-0 px-4 py-2 text-base text-white bg-emerald-300 rounded-lg shadow-md hover:bg-emerald-500 focus:outline-none"
        @click="updateArticle()"
      >
        編集
      </button>
      <button
        class="flex-shrink-0 px-4 py-2 text-base text-white bg-red-300 rounded-lg shadow-md hover:bg-red-500 focus:outline-none"
        @click="deleteArticle()"
      >
        削除
      </button>
    </div>
  </div>
</template>
