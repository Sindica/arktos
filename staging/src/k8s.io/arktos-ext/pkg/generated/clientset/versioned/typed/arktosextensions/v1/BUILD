load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "arktosextensions_client.go",
        "doc.go",
        "generated_expansion.go",
        "network.go",
    ],
    importmap = "k8s.io/kubernetes/vendor/k8s.io/arktos-ext/pkg/generated/clientset/versioned/typed/arktosextensions/v1",
    importpath = "k8s.io/arktos-ext/pkg/generated/clientset/versioned/typed/arktosextensions/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//staging/src/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/watch:go_default_library",
        "//staging/src/k8s.io/arktos-ext/pkg/apis/arktosextensions/v1:go_default_library",
        "//staging/src/k8s.io/arktos-ext/pkg/generated/clientset/versioned/scheme:go_default_library",
        "//staging/src/k8s.io/client-go/apiserverupdate:go_default_library",
        "//staging/src/k8s.io/client-go/rest:go_default_library",
        "//vendor/k8s.io/klog:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//staging/src/k8s.io/arktos-ext/pkg/generated/clientset/versioned/typed/arktosextensions/v1/fake:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
