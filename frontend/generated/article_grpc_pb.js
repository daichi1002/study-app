// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// protoのバージョン
'use strict';
var grpc = require('@grpc/grpc-js');
var article_pb = require('./article_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');

function serialize_article_CreateArticleRequest(arg) {
  if (!(arg instanceof article_pb.CreateArticleRequest)) {
    throw new Error('Expected argument of type article.CreateArticleRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_article_CreateArticleRequest(buffer_arg) {
  return article_pb.CreateArticleRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_article_CreateArticleResponse(arg) {
  if (!(arg instanceof article_pb.CreateArticleResponse)) {
    throw new Error('Expected argument of type article.CreateArticleResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_article_CreateArticleResponse(buffer_arg) {
  return article_pb.CreateArticleResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_article_ListArticlesRequest(arg) {
  if (!(arg instanceof article_pb.ListArticlesRequest)) {
    throw new Error('Expected argument of type article.ListArticlesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_article_ListArticlesRequest(buffer_arg) {
  return article_pb.ListArticlesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_article_ListArticlesResponse(arg) {
  if (!(arg instanceof article_pb.ListArticlesResponse)) {
    throw new Error('Expected argument of type article.ListArticlesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_article_ListArticlesResponse(buffer_arg) {
  return article_pb.ListArticlesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ArticleServiceService = exports.ArticleServiceService = {
  createArticle: {
    path: '/article.ArticleService/CreateArticle',
    requestStream: false,
    responseStream: false,
    requestType: article_pb.CreateArticleRequest,
    responseType: article_pb.CreateArticleResponse,
    requestSerialize: serialize_article_CreateArticleRequest,
    requestDeserialize: deserialize_article_CreateArticleRequest,
    responseSerialize: serialize_article_CreateArticleResponse,
    responseDeserialize: deserialize_article_CreateArticleResponse,
  },
  getArticles: {
    path: '/article.ArticleService/GetArticles',
    requestStream: false,
    responseStream: false,
    requestType: article_pb.ListArticlesRequest,
    responseType: article_pb.ListArticlesResponse,
    requestSerialize: serialize_article_ListArticlesRequest,
    requestDeserialize: deserialize_article_ListArticlesRequest,
    responseSerialize: serialize_article_ListArticlesResponse,
    responseDeserialize: deserialize_article_ListArticlesResponse,
  },
};

exports.ArticleServiceClient = grpc.makeGenericClientConstructor(ArticleServiceService);
