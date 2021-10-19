pragma solidity ^0.8.7;

/**
@dev this is a stripped down version of the ERC721 token standard containing only the methods 
we need to call on the contract code to keep things a minimal as possible
 */
interface ERC721Stripped {
    function setApprovalForAll(address op, bool approved) external;

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId
    ) external;

    function balanceOf(address owner) external view returns (uint256 balance);

    function tokenOfOwnerByIndex(address owner, uint256 index)
        external
        view
        returns (uint256 tokenId);
}

/// @dev Note: the ERC-165 identifier for this interface is 0x150b7a02.
interface ERC721TokenReceiver {
    function onERC721Received(
        address _operator,
        address _from,
        uint256 _tokenId,
        bytes calldata _data
    ) external returns (bytes4);
}

interface dotdotdot is ERC721Stripped {
    function mint(uint256 numberOfTokensMax5) external payable;
}

contract dotdotbot is ERC721TokenReceiver {
    // the address of the dotdotdot contract implementation
    address private _dotdotdotContract;
    address private _owner;
    mapping(address => bool) private whitelisted;

    uint256 private PRICE; // this is the base price of the mint
    uint256 private DOUBLE_PRICE; // This is the minimum balance we want to have in the contract for any potential gas fees that may incur

    constructor() {
        _owner = msg.sender;
        whitelisted[msg.sender] = true; // deployer of the contract gets whitelisted
    }

    function onERC721Received(
        address _operator,
        address _from,
        uint256 _tokenId,
        bytes calldata _data
    ) external override returns (bytes4) {
        return dotdotbot.onERC721Received.selector;
    }

    modifier onlyOwner() {
        require(
            msg.sender == _owner,
            "you must be an owner to execute this function"
        );
        _;
    }

    modifier onlyWhitelist() {
        require(
            whitelisted[msg.sender],
            "you must be whitelisted to perform this action"
        );
        _;
    }

    function setParameters(
        address implementationAddress,
        uint256 nftPrice) external onlyOwner {
        PRICE = nftPrice; // in wei
        DOUBLE_PRICE = nftPrice * 2; // the minimal reserves required for the mint
        _dotdotdotContract = implementationAddress;
    }

    function setWhitelisted(address addr, bool status) public onlyOwner {
        whitelisted[addr] = status;
    }

    function owner() external view returns (address) {
        return _owner;
    }

    function deposit() public payable onlyWhitelist {}

    function withdraw() public onlyOwner {
        payable(_owner).transfer(address(this).balance);
    }

    function tryMint(uint256 count) public onlyWhitelist {
        if (address(this).balance >= DOUBLE_PRICE  * count) {
            dotdotdot(_dotdotdotContract).mint{value: PRICE * count}(count);
        }
    }

    function transferToOwner() public onlyOwner {
        ERC721Stripped dot = ERC721Stripped(_dotdotdotContract);

        dot.setApprovalForAll(_owner, true);
        uint256 balance = dot.balanceOf(address(this));

        for (uint256 i = 0; i < balance; i++) {
            uint256 token = dot.tokenOfOwnerByIndex(address(this), 0);
            dot.safeTransferFrom(address(this), _owner, token);
        }
    }
}
