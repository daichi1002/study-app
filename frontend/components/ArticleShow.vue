<script setup lang="ts">
import { Article } from "~/types/article";
import MdEditor from "md-editor-v3";
import { formatDate, setTagName, setUserName } from "~/util/format";

interface Props {
  value: Article;
}
const props = defineProps<Props>();
</script>

<template>
  <div class="flex text-2xl mt-4">
    <h1 class="mr-10">{{ setUserName(props.value.userId) }}</h1>
    <h1>{{ formatDate(props.value.updatedAt) }}</h1>
  </div>
  <div class="flex text-2xl mt-6">
    <h1 class="mr-10">{{ setTagName(props.value.tags) }}</h1>
  </div>
  <div class="mt-6">
    <md-editor v-model="props.value.content" preview-only />
  </div>
</template>
