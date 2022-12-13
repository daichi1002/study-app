<script setup lang="ts">
import { Article } from "~/types/article";
const { data } = await useFetch<Article>("http://localhost:8080/");
</script>

<template>
  <div class="container mx-auto p-7" v-if="data">
    <div class="flex">
      <h1 class="text-4xl font-bold mr-4">{{ data.title }}</h1>
    </div>
    <ArticleShow :value="data" />
  </div>
</template>
