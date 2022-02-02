import React from "react";
import Showcase from "./Showcase";

const Header = () => {
  return (
    <div className="w-full h-screen flex flex-row">
      <Showcase />
      <div className="bg-gray-100 w-full">Coming soon :D</div>
    </div>
  );
};

export default Header;
