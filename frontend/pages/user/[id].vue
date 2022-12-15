<script setup lang="ts">
import { User } from "~/types/user";
import { formatDate } from "~/util/format";

const router = useRoute();
const { id } = router.params;

const chartOptions = {
  chart: {
    id: "vuechart-example",
  },
  xaxis: {
    categories: [1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998],
  },
};
const series = [
  {
    name: "series-1",
    data: [30, 40, 35, 50, 49, 60, 70, 91],
  },
];

const { data } = await useFetch<User>(`http://localhost:8080/user/${id}`);
</script>
<template>
  <div class="container max-w-5xl px-4 mx-auto sm:px-8">
    <div class="py-8">
      <div class="flex flex-row justify-between w-full mb-1 sm:mb-0">
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
        <div>
          <div>学習している言語</div>
          <ClientOnly>
            <apexchart
              width="500"
              type="bar"
              :options="chartOptions"
              :series="series"
            ></apexchart>
          </ClientOnly>
        </div>
      </div>
    </div>
  </div>
</template>
