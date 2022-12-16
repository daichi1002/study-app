<script setup lang="ts">
import { User } from "~/types/user";
import { formatDate } from "~/util/format";

const router = useRoute();
const { id } = router.params;

const category: Array<string> = [];
const values: Array<number> = [];

const { data } = await useFetch<User>(`http://localhost:8080/user/${id}`);

data.value?.languages.map((language) => {
  category.push(language.name);
  values.push(language.ratio);
});

const chartOptions = {
  chart: {
    id: "vuechart-example",
  },
  xaxis: {
    type: "category",
    categories: category,
  },
  yaxis: {
    max: 100, // 最大値
  },
};
const series = [
  {
    name: "series-1",
    data: values,
  },
];
</script>
<template>
  <div class="container max-w-5xl px-4 mx-auto sm:px-8">
    <div class="py-8">
      <div class="flex flex-row grid grid-cols-3 gap-4 w-full mb-1 sm:mb-0">
        <div class="max-w-sm rounded overflow-hidden shadow-lg">
          <img
            class="w-full"
            src="~/assets/images/kamoshika.png"
            alt="Sunset in the mountains"
          />
          <div class="px-6 py-4">
            <div class="font-bold text-xl mb-2">{{ data?.userName }}さん</div>
            <div class="flex text-gray-700 text-base">
              <p>生年月日：</p>
              <p>
                {{ formatDate(data?.birthday) }}
              </p>
            </div>
            <div class="flex text-gray-700 text-base">
              <p>所属PRJ：</p>
              <p>
                {{ data?.affiliation }}
              </p>
            </div>
          </div>
        </div>
        <div
          class="col-span-2 rounded content-center overflow-hidden shadow-lg"
        >
          <div class="p-4 font-bold text-xl">学習している言語</div>
          <ClientOnly>
            <apexchart
              height="300"
              type="bar"
              :options="chartOptions"
              :series="series"
            ></apexchart>
          </ClientOnly>
        </div>
      </div>
      <div class="py-8">
        <div class="flex flex-row justify-between w-full mb-1 sm:mb-0">
          <h2 class="text-2xl leading-tight">投稿記事</h2>
        </div>
        <div class="px-4 py-4 -mx-4 overflow-x-auto sm:-mx-8 sm:px-8">
          <div
            class="inline-block min-w-full overflow-hidden rounded-lg shadow-md"
          >
            <table class="min-w-full leading-normal">
              <thead>
                <tr>
                  <th
                    scope="col"
                    class="px-5 py-3 text-sm font-normal text-left text-gray-800 uppercase bg-white border-b border-gray-300"
                  >
                    No.
                  </th>
                  <th
                    scope="col"
                    class="px-5 py-3 text-sm font-normal text-left text-gray-800 uppercase bg-white border-b border-gray-300"
                  >
                    タイトル
                  </th>
                  <th
                    scope="col"
                    class="px-5 py-3 text-sm font-normal text-left text-gray-800 uppercase bg-white border-b border-gray-300"
                  >
                    更新日
                  </th>
                  <th
                    scope="col"
                    class="px-5 py-3 text-sm font-normal text-left text-gray-800 uppercase bg-white border-b border-gray-200"
                  ></th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="(article, i) in data?.articles"
                  class="hover:bg-slate-100"
                >
                  <td
                    class="px-5 py-5 text-sm bg-white border-b border-gray-300"
                  >
                    {{ i + 1 }}
                  </td>
                  <td
                    class="px-5 py-5 text-sm bg-white border-b border-gray-300"
                  >
                    {{ article.title }}
                  </td>
                  <td
                    class="px-5 py-5 text-sm bg-white border-b border-gray-300"
                  >
                    {{ formatDate(article.updatedAt) }}
                  </td>
                  <td
                    class="px-5 py-5 text-sm bg-white border-b border-gray-300"
                  >
                    <NuxtLink :to="`/article/${article.articleId}`"
                      ><button
                        class="flex-shrink-0 px-4 py-2 text-base text-white bg-emerald-300 rounded-lg shadow-md hover:bg-emerald-500 focus:outline-none"
                        type="submit"
                      >
                        詳細
                      </button></NuxtLink
                    >
                  </td>
                </tr>
              </tbody>
            </table>
            <div
              class="flex flex-col items-center px-5 py-5 bg-white xs:flex-row xs:justify-between"
            >
              <div class="flex items-center">
                <button
                  type="button"
                  class="w-full p-4 text-base text-gray-600 bg-white border rounded-l-xl hover:bg-gray-100"
                >
                  <svg
                    width="9"
                    fill="currentColor"
                    height="8"
                    class=""
                    viewBox="0 0 1792 1792"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      d="M1427 301l-531 531 531 531q19 19 19 45t-19 45l-166 166q-19 19-45 19t-45-19l-742-742q-19-19-19-45t19-45l742-742q19-19 45-19t45 19l166 166q19 19 19 45t-19 45z"
                    ></path>
                  </svg>
                </button>
                <button
                  type="button"
                  class="w-full px-4 py-2 text-base text-indigo-500 bg-white border-t border-b hover:bg-gray-100"
                >
                  1
                </button>
                <button
                  type="button"
                  class="w-full px-4 py-2 text-base text-gray-600 bg-white border hover:bg-gray-100"
                >
                  2
                </button>
                <button
                  type="button"
                  class="w-full px-4 py-2 text-base text-gray-600 bg-white border-t border-b hover:bg-gray-100"
                >
                  3
                </button>
                <button
                  type="button"
                  class="w-full px-4 py-2 text-base text-gray-600 bg-white border hover:bg-gray-100"
                >
                  4
                </button>
                <button
                  type="button"
                  class="w-full p-4 text-base text-gray-600 bg-white border-t border-b border-r rounded-r-xl hover:bg-gray-100"
                >
                  <svg
                    width="9"
                    fill="currentColor"
                    height="8"
                    class=""
                    viewBox="0 0 1792 1792"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      d="M1363 877l-742 742q-19 19-45 19t-45-19l-166-166q-19-19-19-45t19-45l531-531-531-531q-19-19-19-45t19-45l166-166q19-19 45-19t45 19l742 742q19 19 19 45t-19 45z"
                    ></path>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
