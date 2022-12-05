<script setup lang="ts">
import { Article } from "~/types/article";
const router = useRoute();
const { id } = router.params;

const { data } = await useFetch<Article>(`http://localhost:8080/article/${id}`);
</script>

<template>
  <div class="container mx-auto p-7" v-if="data">
    <div class="flex">
      <h1 class="text-4xl font-bold mr-4">{{ data.title }}</h1>
      <NuxtLink to="">
        <button class="w-16 h-10 py-2 px-4 rounded-lg">
          <img src="../../assets/images/edit.png" />
        </button>
      </NuxtLink>
    </div>
    <ArticleShow :value="data" />
  </div>
</template>
