import React from "react";
import logo from "../images/logo.svg";

const Header = () => {
  return (
    <div className="w-full h-screen flex flex-row">
      <div className="w-full flex flex-row justify-center bg-purple-900">
        <div className="max-w-2xl w-11/12 h-full flex flex-col justify-center items-center">
          <img
            // // className="max-h-[50%] h-full" // Forgive me, for i have to commit some evilness.
            className="max-h-full filter invert"
            src={logo}
          />
          <button className="mt-8 border-2 border-gray-50 px-6 pt-2 pb-3 text-xl font-bold text-white rounded-full">
            Project Showcase
          </button>
        </div>
      </div>
      <div className="bg-gray-100 w-full">yes</div>
    </div>
  );
};

export default Header;
