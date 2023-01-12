# WooFi

This connector ingests the woofi contracts specified below. Related connectors are in ../woofi_avax, ../woofi_arbi, ../woofi_bsc, ../woofi_fantom, ../woofi_polygon.

ABIs are based on the ones [provided by the Woofi Team](https://github.com/woonetwork/woofi_subgraph/tree/main/abis), with their names changed slightly to make it backwards compatible with the currently deployed connector + its users.

## Address <> ABI mappings | start

-- Avalanche --

0xF8cE0D043891b62c55380fB1EFBfB4F186153D96 |  9713408 | WooPPV1_1               | ABI: WooPPV3

0x1df3009c57a8B143c6246149F00B090Bce3b8f88 | 13965201 | WooPPV1_2               | ABI: WooPPV3

0x3b3E4b4741e91aF52d0e9ad8660573E951c88524 | 23878037 | WooPPV2_1               | ABI: WooPPV4

0x160020B09DeD3d862f7f851B5c50632BcF2062FF | 10857815 | WooRouterV1_1           | ABI: WooRouterV2

0x3e0Da0A9e4139B32b37710784b8DCa643c152001 | 13522432 | WooRouterV1_2           | ABI: WooRouterV2

0x5AA6a4E96A9129562e2fc06660D07FEdDAAf7854 | 13966541 | WooRouterV1_3           | ABI: WooRouterV2

0xC22FBb3133dF781E6C25ea6acebe2D2Bb8CeA2f9 | 23878146 | WooRouterV2_1           | ABI: WooRouterV3

0xdF37F7A85D4563f39A78494568824b4dF8669B7a | 16011735 | WooCrossChainRouterV1_1 | ABI: WooCrossChainRouterV1

0x1E6bB552ac038c6AFB6EC5Db6B06fDd106e31e33 | 23878188 | WooCrossChainRouterV1_2 | ABI: WooCrossChainRouterV1

-- Arbitrum --

0xeFF23B4bE1091b53205E35f3AfCD9C7182bf3062 | 36279764 | WooPPV2_2               | ABI: WooPPV4

0x37a9dE70b6734dFCA54395D8061d9411D9910739 | 34401135 | WooracleV2              | ABI: WooracleV2

0x9aEd3A8896A85FE9a8CAc52C9B402D092B629a30 | 35134665 | WooRouterV2_1           | ABI: WooRouterV3

0x0972A0fa37984E7FF2aEFA53A0Bb10dCE535aa73 | 34887830 | WooCrossChainRouterV1_1 | ABI: WooCrossChainRouterV1

-- BSC --

0x10C24658815585851a8888f059Cb4338790023F1 | 11981813 | WooPPV1_0               | ABI: WooPPV1

0x8489d142Da126F4Ea01750e80ccAa12FD1642988 | 12094867 | WooPPV1_1               | ABI: WooPPV2

0xbf365Ce9cFcb2d5855521985E351bA3bcf77FD3F | 13662625 | WooPPV1_2               | ABI: WooPPV3

0xEc054126922a9a1918435c9072c32f1B60cB2B90 | 24066408 | WooPPV2_1               | ABI: WooPPV4

0x114f84658c99aa6EA62e3160a87A16dEaF7EFe83 | 11983352 | WooRouterV1_1           | ABI: WooRouterV1

0xcEf5BE73ae943B77f9Bc08859367D923C030a269 | 17015221 | WooRouterV1_2           | ABI: WooRouterV2

0xC90bFE9951a4Efbf20aCa5ECd9966b2bF8A01294 | 24066455 | WooRouterV2_1           | ABI: WooRouterV3

0x53E255e8Bbf4EDF16797f9885291B3Ca0C70B59f | 18675185 | WooCrossChainRouterV1_1 | ABI: WooCrossChainRouterV1

0xd12D239b781e34E0AAa106159940803A07E31a67 | 24066491 | WooCrossChainRouterV1_2 | ABI: WooCrossChainRouterV1

-- Fantom --

0x9503E7517D3C5bc4f9E4A1c6AE4f8B33AC2546f2 | 35475008 | WooPPV1_1               | ABI: WooPPV3

0x286ab107c5E9083dBed35A2B5fb0242538F4f9bf | 52583548 | WooPPV2_1               | ABI: WooPPV4

0x37B5a5A730dAD670874f26Cc5507bb1b9705e447 | 35476301 | WooRouterV1_1           | ABI: WooRouterV2

0x382A9b0bC5D29e96c3a0b81cE9c64d6C8F150Efb | 52583652 | WooRouterV2_1           | ABI: WooRouterV3

0xcF6Ce5Fd6bf28bB1AeAc88A55251f6c840059De5 | 40483282 | WooCrossChainRouterV1_1 | ABI: WooCrossChainRouterV1

0x28D2B949024FE50627f1EbC5f0Ca3Ca721148E40 | 52583699 | WooCrossChainRouterV1_2 | ABI: WooCrossChainRouterV1

-- Polygon --

0x7400B665C8f4f3a951a99f1ee9872efb8778723d | 29860023 | WooPPV1_1               | ABI: WooPPV3

0x7081A38158BD050Ae4a86e38E0225Bc281887d7E | 36856771 | WooPPV2_1               | ABI: WooPPV4

0x9D1A92e601db0901e69bd810029F2C14bCCA3128 | 29860251 | WooRouterV1_1           | ABI: WooRouterV2

0x817Eb46D60762442Da3D931Ff51a30334CA39B74 | 36859029 | WooRouterV2_1           | ABI: WooRouterV3

0x376d567C5794cfc64C74852A9DB2105E0b5B482C | 29897168 | WooCrossChainRouterV1_1 | ABI: WooCrossChainRouterV1

0x574b9cec19553435B360803D8B4De2a5b2C008Fd | 36882842 | WooCrossChainRouterV1_2 | ABI: WooCrossChainRouterV1

## Naming scheme

The ABI names are based on backwards compatibility and contract names.

subgraph/WooPPV1_0.json               => WooPPV1.abi

subgraph/WooPPV1_1.json               => WooPPV2.abi

subgraph/WooPPV1_2.json               => WooPPV3.abi

subgraph/WooPPV2_1.json               => WooPPV4.abi

subgraph/WooRouterV1_1.json           => WooRouterV1.abi

subgraph/WooRouterV1_2.json           => WooRouterV2.abi

subgraph/WooRouterV2_1.json           => WooRouterV3.abi

subgraph/WooCrossChainRouterV1_1.json => WooCrossChainRouterV1.abi
