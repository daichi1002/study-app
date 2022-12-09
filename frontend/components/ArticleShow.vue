<script setup lang="ts">
import { Article, ArticleTag } from "~/types/article";
import MdEditor from "md-editor-v3";
import { formatDate } from "~/util/date";

interface Props {
  value: Article;
}
const props = defineProps<Props>();

const setTagName = (tags: ArticleTag[]) => {
  let name = "";
  const len = tags.length;
  tags.forEach((tag, i) => {
    if (len == i + 1) {
      name += `${tag.tagName}`;
      return;
    }
    name += `${tag.tagName} / `;
  });
  return name;
};
</script>

<template>
  <div class="flex text-2xl mt-4">
    <h1 class="mr-10">{{ props.value.userId }}</h1>
    <h1>{{ formatDate(props.value.updatedAt) }}</h1>
  </div>
  <div class="flex text-2xl mt-6">
    <h1 class="mr-10">{{ setTagName(props.value.tags) }}</h1>
  </div>
  <div class="mt-6">
    <md-editor v-model="props.value.content" preview-only />
  </div>
</template>
