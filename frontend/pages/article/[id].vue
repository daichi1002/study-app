<script setup lang="ts">
import { Article } from "~/types/article";
const router = useRoute();
const { id } = router.params;

const { data } = await useFetch<Article>(`http://localhost:8080/article/${id}`);
</script>

<template>
  <div class="container mx-auto p-7" v-if="data">
    <ArticleShow :value="data" />
  </div>
</template>
