<script setup lang="ts">
import { Tag } from "~/types/tag";
import { useField } from "vee-validate";
import * as yup from "yup";

const { data: tags, error } = await useFetch<Tag[]>(
  `http://localhost:8080/tags`
);

if (!tags.value || error.value) {
  throw createError({
    statusCode: 404,
    message: "Tags not found",
  });
}

const { article, setTitle, setTags, setContent } = useArticle();
const title = ref(article.value.title);
const content = ref(article.value.content);

// const { errorMessage: titleError, value: inputTitle } = useField(
//   ref(article.value.title),
//   yup
//     .string()
//     .required("入力必須項目です")
//     .max(30, "30文字以内で入力してください")
// );
// const { errorMessage: contentError, value: inputContent } = useField(
//   ref(article.value.content),
//   yup.string().required("入力必須項目です")
// );
</script>

<template>
  <div class="mt-5">
    <div class="flex">
      <div class="mr-8">タイトル</div>
      <!-- <div class="text-red-500">{{ titleError }}</div> -->
    </div>
    <input
      class="appearance-none border-2 border-gray-300 rounded w-2/4 py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
      id="inline-full-name"
      type="text"
      v-model="title"
      @blur="setTitle(title)"
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
          @blur="setTags(tag)"
          class="w-4 h-4 text-blue-600 bg-gray-100 rounded border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
        />
        <label for="default-checkbox" class="ml-2 text-sm font-medium">{{
          tag.tagName
        }}</label>
      </div>
    </div>
  </div>
  <div class="mt-4">
    <div class="flex">
      <div class="mr-8">本文</div>
      <!-- <div class="text-red-500">{{ contentError }}</div> -->
    </div>
    <Markdown v-model="content" @onChange="setContent(content)" />
  </div>
</template>
