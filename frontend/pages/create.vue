<script setup lang="ts">
import { Tag } from "~/types/tag";

const { data: tags } = await useFetch<Tag[]>(`http://localhost:8080/create`);
</script>

<template>
  <div class="container mx-auto">
    <div class="mt-10">
      <div>タイトル</div>
      <input
        class="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-2/4 py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
        id="inline-full-name"
        type="text"
        value=""
      />
    </div>
    <div class="mt-10">
      <div>タグ</div>
      <div class="grid grid-cols-5 gap-5">
        <div v-for="tag in tags">
          <input
            id="default-checkbox"
            type="checkbox"
            value=""
            class="w-4 h-4 text-blue-600 bg-gray-100 rounded border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
          />
          <label for="default-checkbox" class="ml-2 text-sm font-medium">{{
            tag.name
          }}</label>
        </div>
      </div>
    </div>
    <div class="mt-10">
      <div>本文</div>
      <textarea
        class="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full h-72 py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
        id="inline-full-name"
        type="text"
        value=""
      />
    </div>
    <div class="mt-6 grid grid-cols-9 gap-9">
      <button class="bg-teal-300 hover:bg-teal-500 py-2 px-4 rounded-lg">
        保存
      </button>
      <button class="bg-red-400 hover:bg-red-500 py-2 px-4 rounded-lg">
        削除
      </button>
    </div>
  </div>
</template>
