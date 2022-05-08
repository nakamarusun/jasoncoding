import React from "react";

const showcase = () => {
  return (
    <div className="flex flex-col">
      <div className="flex flex-row justify-center h-screen bg-neutral-900">
        <div className="max-w-5xl w-11/12 h-full flex flex-col">
          <div className="flex-1">yes</div>
          <div className="w-full bg-gray-600 h-[1px]" />
          <h1 className="flex-initial my-4 text-center text-purple-500 font-black text-xl flex flex-row justify-center">
            <span className="material-icons text-center text-2xl animate-bounce">
              arrow_drop_down
            </span>
            <span>Scroll and Enjoy</span>
            <span className="align-super text-sm">0</span>
          </h1>
        </div>
      </div>
    </div>
  );
};

export default showcase;
