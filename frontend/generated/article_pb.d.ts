// package: article
// file: article.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Article extends jspb.Message {
  getArticleId(): string;
  setArticleId(value: string): void;

  getUserId(): string;
  setUserId(value: string): void;

  getTitle(): string;
  setTitle(value: string): void;

  getContent(): string;
  setContent(value: string): void;

  hasCreatedAt(): boolean;
  clearCreatedAt(): void;
  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;
  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Article.AsObject;
  static toObject(includeInstance: boolean, msg: Article): Article.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Article, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Article;
  static deserializeBinaryFromReader(message: Article, reader: jspb.BinaryReader): Article;
}

export namespace Article {
  export type AsObject = {
    articleId: string,
    userId: string,
    title: string,
    content: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class ListArticlesRequest extends jspb.Message {
  getPage(): number;
  setPage(value: number): void;

  getPageSize(): number;
  setPageSize(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListArticlesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListArticlesRequest): ListArticlesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListArticlesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListArticlesRequest;
  static deserializeBinaryFromReader(message: ListArticlesRequest, reader: jspb.BinaryReader): ListArticlesRequest;
}

export namespace ListArticlesRequest {
  export type AsObject = {
    page: number,
    pageSize: number,
  }
}

export class ListArticlesResponse extends jspb.Message {
  clearArticlesList(): void;
  getArticlesList(): Array<Article>;
  setArticlesList(value: Array<Article>): void;
  addArticles(value?: Article, index?: number): Article;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListArticlesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListArticlesResponse): ListArticlesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListArticlesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListArticlesResponse;
  static deserializeBinaryFromReader(message: ListArticlesResponse, reader: jspb.BinaryReader): ListArticlesResponse;
}

export namespace ListArticlesResponse {
  export type AsObject = {
    articlesList: Array<Article.AsObject>,
  }
}

export class CreateArticleRequest extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): void;

  clearTagIdsList(): void;
  getTagIdsList(): Array<string>;
  setTagIdsList(value: Array<string>): void;
  addTagIds(value: string, index?: number): string;

  getTitle(): string;
  setTitle(value: string): void;

  getContent(): string;
  setContent(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateArticleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateArticleRequest): CreateArticleRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateArticleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateArticleRequest;
  static deserializeBinaryFromReader(message: CreateArticleRequest, reader: jspb.BinaryReader): CreateArticleRequest;
}

export namespace CreateArticleRequest {
  export type AsObject = {
    userId: string,
    tagIdsList: Array<string>,
    title: string,
    content: string,
  }
}

export class CreateArticleResponse extends jspb.Message {
  hasArticle(): boolean;
  clearArticle(): void;
  getArticle(): Article | undefined;
  setArticle(value?: Article): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateArticleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateArticleResponse): CreateArticleResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateArticleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateArticleResponse;
  static deserializeBinaryFromReader(message: CreateArticleResponse, reader: jspb.BinaryReader): CreateArticleResponse;
}

export namespace CreateArticleResponse {
  export type AsObject = {
    article?: Article.AsObject,
  }
}

