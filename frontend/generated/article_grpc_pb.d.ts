// GENERATED CODE -- DO NOT EDIT!

// package: article
// file: article.proto

import * as article_pb from "./article_pb";
import * as grpc from "@grpc/grpc-js";

interface IArticleServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  createArticle: grpc.MethodDefinition<article_pb.CreateArticleRequest, article_pb.CreateArticleResponse>;
  getArticles: grpc.MethodDefinition<article_pb.ListArticlesRequest, article_pb.ListArticlesResponse>;
}

export const ArticleServiceService: IArticleServiceService;

export interface IArticleServiceServer extends grpc.UntypedServiceImplementation {
  createArticle: grpc.handleUnaryCall<article_pb.CreateArticleRequest, article_pb.CreateArticleResponse>;
  getArticles: grpc.handleUnaryCall<article_pb.ListArticlesRequest, article_pb.ListArticlesResponse>;
}

export class ArticleServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  createArticle(argument: article_pb.CreateArticleRequest, callback: grpc.requestCallback<article_pb.CreateArticleResponse>): grpc.ClientUnaryCall;
  createArticle(argument: article_pb.CreateArticleRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<article_pb.CreateArticleResponse>): grpc.ClientUnaryCall;
  createArticle(argument: article_pb.CreateArticleRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<article_pb.CreateArticleResponse>): grpc.ClientUnaryCall;
  getArticles(argument: article_pb.ListArticlesRequest, callback: grpc.requestCallback<article_pb.ListArticlesResponse>): grpc.ClientUnaryCall;
  getArticles(argument: article_pb.ListArticlesRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<article_pb.ListArticlesResponse>): grpc.ClientUnaryCall;
  getArticles(argument: article_pb.ListArticlesRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<article_pb.ListArticlesResponse>): grpc.ClientUnaryCall;
}
