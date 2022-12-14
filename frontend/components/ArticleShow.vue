<script setup lang="ts">
import { Article } from "~/types/article";
import MdEditor from "md-editor-v3";
import { formatDate, setTagName, setUserName } from "~/util/format";
import { User } from "~/types/user";

interface Props {
  value: Article;
}
const props = defineProps<Props>();
const { data: users } = await useFetch<User[]>("http://localhost:8080/users");
</script>

<template>
  <div class="flex text-lg mt-2">
    <h1 class="mr-10">{{ setUserName(props.value.userId, users) }}</h1>
    <h1>{{ formatDate(props.value.updatedAt) }}</h1>
  </div>
  <div class="flex text-lg mt-3 grid grid-cols-12">
    <div v-for="tag in props.value.tags">
      <div class="mr-2 border rounded-lg border-emerald-300 text-center">
        {{ tag.tagName }}
      </div>
    </div>
  </div>
  <div class="mt-6">
    <md-editor v-model="props.value.content" preview-only />
  </div>
</template>
