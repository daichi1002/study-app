import { viteCommonjs } from "@originjs/vite-plugin-commonjs";

export default {
  optimizeDeps: {
    include: ["generated/article_grpc_pb"],
  },
};
