<template>
  <header class="border text-gray-600 body-font">
    <div
      class="container mx-auto flex flex-wrap p-4 flex-col md:flex-row items-center"
    >
      <NuxtLink to="/"><img src="../assets/images/study.png" /></NuxtLink>
      <nav
        class="md:mr-auto md:ml-4 md:py-1 md:pl-4 md:border-l md:border-gray-400 flex flex-wrap items-center text-base justify-center"
      >
        <NuxtLink to="/article/list" class="mr-5 hover:text-gray-900"
          >記事</NuxtLink
        >
        <NuxtLink to="/article/list" class="mr-5 hover:text-gray-900"
          >チーム分け</NuxtLink
        >
        <NuxtLink to="/article/list" class="mr-5 hover:text-gray-900"
          >ユーザー</NuxtLink
        >
      </nav>
      <button
        class="inline-flex items-center bg-gray-100 border-0 py-1 px-3 focus:outline-none hover:bg-gray-200 rounded text-base mt-4 md:mt-0"
      >
        Button
        <svg
          fill="none"
          stroke="currentColor"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          class="w-4 h-4 ml-1"
          viewBox="0 0 24 24"
        >
          <path d="M5 12h14M12 5l7 7-7 7"></path>
        </svg>
      </button>
    </div>
  </header>
</template>
