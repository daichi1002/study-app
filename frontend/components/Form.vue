<script setup lang="ts">
import { Tag } from "~/types/tag";
import { Article } from "~/types/article";

interface Props {
  value: Article;
}

const { data: tags, error } = await useFetch<Tag[]>(
  `http://localhost:8080/tags`
);

if (!tags.value || error.value) {
  throw createError({
    statusCode: 404,
    message: "Tags not found",
  });
}

const props = defineProps<Props>();
</script>

<template>
  <div class="mt-5">
    <div>タイトル</div>
    <input
      class="appearance-none border-2 border-gray-300 rounded w-2/4 py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
      id="inline-full-name"
      type="text"
      v-model="props.value.title"
    />
  </div>
  <div class="mt-4">
    <div>タグ</div>
    <div class="grid grid-cols-5 gap-5">
      <div v-for="tag in tags">
        <input
          id="default-checkbox"
          type="checkbox"
          :value="tag"
          v-model="props.value.tags"
          class="w-4 h-4 text-blue-600 bg-gray-100 rounded border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
        />
        <label for="default-checkbox" class="ml-2 text-sm font-medium">{{
          tag.tagName
        }}</label>
      </div>
    </div>
  </div>
  <div class="mt-4">
    <div>本文</div>
    <Markdown v-model="props.value.content" />
  </div>
</template>
