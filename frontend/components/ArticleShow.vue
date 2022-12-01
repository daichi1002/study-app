<script setup lang="ts">
import { Article } from "~/types/article";

interface Props {
  value: Article;
}
const props = defineProps<Props>();
</script>

<template>
  <!-- <div class="container mx-auto p-7">
    <h1 class="text-4xl font-bold">{{ props.value.title }}</h1>
    <div class="flex text-2xl mt-4">
      <h1 class="mr-10">{{ props.value.userName }}</h1>
      <h1>{{ props.value.updatedDate }}</h1>
    </div>
    <div class="flex text-2xl mt-6" v-for="tag in props.value.tag">
      <h1 class="mr-10">{{ tag }}</h1>
    </div>
    <div class="mt-6">
      {{ props.value.content }}
    </div>
  </div> -->
  <h1 class="text-4xl font-bold">Figmaの使い方について</h1>
  <div class="flex text-2xl mt-4">
    <h1 class="mr-10">こじゃ</h1>
    <h1>2022/08/10</h1>
  </div>
  <div class="flex text-2xl mt-6">
    <h1 class="mr-10">デザイン</h1>
    <h1>フロント</h1>
  </div>
  <div class="mt-6">
    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
    tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
    quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
    consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
    cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat
    non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
    tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
    quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
    consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
    cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat
    non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
  </div>
</template>
