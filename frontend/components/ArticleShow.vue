<script setup lang="ts">
import { Article } from "~/types/article";
import MdEditor from "md-editor-v3";

interface Props {
  value: Article;
}
const props = defineProps<Props>();
</script>

<template>
  <div class="flex text-2xl mt-4">
    <h1 class="mr-10">{{ props.value.userName }}</h1>
    <h1>{{ props.value.updatedDate }}</h1>
  </div>
  <div class="flex text-2xl mt-6">
    <h1 class="mr-10">{{ props.value.tag }}</h1>
  </div>
  <div class="mt-6">
    <md-editor v-model="props.value.content" preview-only />
  </div>
</template>
