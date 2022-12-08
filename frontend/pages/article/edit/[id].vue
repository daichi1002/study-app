<script setup lang="ts">
import { Article } from "~/types/article";
import { changePage } from "~/util/router";

const router = useRoute();
const { id } = router.params;

const { data } = await useFetch<Article>(`http://localhost:8080/article/${id}`);

const { article, setTitle, setTags, setContent } = useArticle();

const updateArticle = async () => {
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
  <div class="container mx-auto" v-if="article">
    <Form />
    <div class="mt-4 grid grid-cols-9 gap-9">
      <button
        class="bg-teal-300 hover:bg-teal-500 py-2 px-4 rounded-lg"
        @click="updateArticle()"
      >
        編集
      </button>
      <button
        class="bg-red-400 hover:bg-red-500 py-2 px-4 rounded-lg"
        @click="deleteArticle()"
      >
        削除
      </button>
    </div>
  </div>
</template>
