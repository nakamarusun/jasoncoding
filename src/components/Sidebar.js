import React, { useState } from "react";
import PropTypes from "prop-types";
import { menus } from "../menu";

const Sidebar = ({ className }) => {
  const [open, setOpen] = useState(false);

  return (
    <>
      <div
        className={`${className} bg-black fixed top-0 inset-x-0 h-12 z-40 flex flex-row items-center md:hidden`}
      >
        <div
          className="material-icons text-white ml-4 cursor-pointer"
          onClick={() => {
            setOpen(!open);
          }}
        >
          {open ? "close" : "menu"}
        </div>
      </div>
      {open && (
        <div
          className="md:hidden inset-0 bg-black bg-opacity-50 fixed z-20"
          onClick={() => {
            setOpen(false);
          }}
        />
      )}
      <div
        className={`md:hidden flex flex-col overflow-x-hidden whitespace-nowrap fixed inset-y-0 mt-12 z-30 bottom-0 left-0 bg-gray-900 transition-[width] ${
          open ? "w-48" : "w-0"
        }`}
      >
        {menus.map(({ name, url, icon }, i) => (
          <a className="text-white mt-4 ml-2" key={i} href={url}>
            <span className="material-icons align-bottom text-white mr-4">
              {icon}
            </span>
            <span className="underline decoration-transparent active:decoration-white font-bold">
              {name}
            </span>
          </a>
        ))}
      </div>
    </>
  );
};

Sidebar.propTypes = {
  className: PropTypes.string,
};

export default Sidebar;
