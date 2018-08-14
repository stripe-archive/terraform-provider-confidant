load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/stripe/terraform-provider-confidant",
)

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        "provider.go",
        "resource_credential_assignment.go",
        "resource_service.go",
    ],
    importpath = "github.com/stripe/terraform-provider-confidant",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_hashicorp_terraform//helper/schema:go_default_library",
        "@com_github_hashicorp_terraform//plugin:go_default_library",
        "@com_github_hashicorp_terraform//terraform:go_default_library",
        "@com_github_stripe_go_confidant_client//confidant:go_default_library",
        "@com_github_stripe_go_confidant_client//kmsauth:go_default_library",
    ],
)

go_binary(
    name = "terraform-provider-confidant",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
