import PackageDescription

let package = Package(
    name: "CharonClient",
    products: [
        .library(
            name: "CharonClient",
            targets: ["CharonClient"]
        )
    ],
    dependencies: [
        .package(url: "https://github.com/apple/swift-protobuf.git", from: "1.28.1"),
        .package(url: "https://github.com/grpc/grpc-swift.git", from: "1.23.1"),
    ],
    targets: [
        .target(
            name: "CharonClient",
            dependencies: [
                .product(name: "SwiftProtobuf", package: "swift-protobuf"),
                .product(name: "GRPC", package: "grpc-swift"),
            ],
            path: "Sources/CharonClient",
            sources: ["Generated"], // Only include the generated Swift files.
            publicHeadersPath: ""
        ),
        .testTarget(
            name: "CharonClientTests",
            dependencies: ["CharonClient"],
            path: "Tests/CharonClientTests"
        )
    ]
)
