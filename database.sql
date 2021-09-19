CREATE TABLE "EmailNotification"(
                                    "id" UUID DEFAULT uuid_generate_v4 () NOT NULL,
                                    "body" VARCHAR(255) NOT NULL,
                                    "from" VARCHAR(255) NOT NULL,
                                    "to"   text[] NOT NULL,
                                    "subject" VARCHAR(255) NOT NULL,
                                    "status" VARCHAR(255) NOT NULL,
                                    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
                                    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "PushedNotification"(
                                     "id" UUID DEFAULT uuid_generate_v4 () NOT NULL,
                                     "api_key" VARCHAR(255) NOT NULL,
                                     "token" text[] NOT NULL,
                                     "Title" VARCHAR(255) NOT NULL,
                                     "body" VARCHAR(255) NOT NULL,
                                     "data" VARCHAR(255) NOT NULL,
                                     "status" VARCHAR(255) NOT NULL,
                                     "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
                                     "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE TABLE "sms"(
                                     "id" UUID DEFAULT uuid_generate_v4 () NOT NULL,
                                     "password" VARCHAR(255) NOT NULL,
                                     "user" VARCHAR(255) NOT NULL,
                                     "sender_id" VARCHAR(255) NOT NULL,
                                     "api_gate_way" VARCHAR(255) NOT NULL,
                                     "call_back_url" VARCHAR(255) NOT NULL,
                                     "body" VARCHAR(255) NOT NULL,
                                     "receiver_phone" VARCHAR(255) NOT NULL,
                                     "status" VARCHAR(255) NOT NULL,
                                     "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
                                     "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);