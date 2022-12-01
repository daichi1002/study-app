<script setup lang="ts">
import { Article } from "~/types/article";
const router = useRoute();
const id = router.params;

const { data } = await useFetch<Article>(`/article/${id}`);
console.log(data);
</script>

<template>
  <div class="container mx-auto p-7" v-if="data">
    <ArticleShow :value="data" />
  </div>
</template>
