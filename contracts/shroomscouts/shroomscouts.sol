pragma solidity ^0.8.0;

interface ShroomScouts {
        function balanceOf(address _owner, uint256 _id) external view returns (uint256);
        function mintToken(uint256 amount) external;
}