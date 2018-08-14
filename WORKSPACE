BAZEL_VERSION = "0.13.0"

BAZEL_LINUX_INSTALLER_SHA = "c90ed6d8478fd543d936702d2eb3ed034f46b2223fac790598db70c161552418"

BAZEL_DARWIN_INSTALLER_SHA = "f7ad77c509eda9bc6cd3c98945858e278265a9551a1d7b31f38558a296b71027"

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.12.1/rules_go-0.12.1.tar.gz"],
    sha256 = "8b68d0630d63d95dacc0016c3bb4b76154fe34fca93efd65d1c366de3fcb4294",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.12.0/bazel-gazelle-0.12.0.tar.gz"],
    sha256 = "ddedc7aaeb61f2654d7d7d4fd7940052ea992ccdb031b8f9797ed143ac7e8d43",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_hashicorp_terraform",
    commit = "2487af19453a0d55a428fb17150f87b24170ccc1",
    importpath = "github.com/hashicorp/terraform",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    commit = "2a14182c3ceee916649d54eb2c16ebe8e57ee326",
    importpath = "github.com/aws/aws-sdk-go",
)

go_repository(
    name = "com_github_stripe_go_confidant_client",
    commit = "d5b23f2f0eb2247b8d10a5176c62f97681c33631",
    importpath = "github.com/stripe/go-confidant-client",
)
