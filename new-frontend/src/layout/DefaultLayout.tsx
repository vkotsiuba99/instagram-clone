import React, { FunctionComponent } from 'react';
import Avatar from '../components/Avatar';
import Icon from '../components/Icon';
import Sidebar from '../components/sidebar/Sidebar';

type DefaultLayoutProps = {};

const DefaultLayout: FunctionComponent<DefaultLayoutProps> = (
  props,
): JSX.Element => {
  return (
    <div className="grid h-screen grid-cols-4 xl:grid-cols-7 2xl:grid-cols-5">
      <Sidebar>
        <div className="px-8 space-y-3">
          <Avatar />
          <div className="space-y-0">
            <h2 className="text-lg font-bold text-center text-gray-800">
              Kate Lingard
            </h2>
            <p className="text-sm text-center text-gray-400">@katy69</p>
          </div>
        </div>

        <div className="grid justify-between grid-cols-3 gap-6 px-8 my-8 text-sm">
          <div className="text-center">
            <p className="font-bold text-gray-800">46</p>
            <p className="text-gray-400">Posts</p>
          </div>
          <div className="text-center">
            <p className="font-bold text-gray-800">2.8k</p>
            <p className="text-gray-400">Followers</p>
          </div>
          <div className="text-center">
            <p className="font-bold text-gray-800">526</p>
            <p className="text-gray-400">Following</p>
          </div>
        </div>

        {/* TODO: Add sidebar items */}

        <div className="flex items-center">
          <div className="flex-grow bg-gray-200 border"></div>
          <div className="p-1 -mr-2 bg-white rounded-full">
            <Icon name="chevron-left" className="w-5 h-5" />
          </div>
        </div>

        <div className="flex items-center pl-12 space-x-6 text-gray-600">
          <Icon name="logout" className="w-5 h-5" />
          <p>Logout</p>
        </div>
      </Sidebar>

      <div className="container col-span-3 px-16 mx-auto my-12 xl:col-span-5 2xl:col-span-4">
        {props.children}
      </div>
    </div>
  );
};

export default DefaultLayout;