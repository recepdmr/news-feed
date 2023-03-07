import React, {PropsWithChildren} from 'react';
import {SafeAreaView} from 'react-native-safe-area-context';

const BaseLayout: React.FC<PropsWithChildren> = ({
  children,
}: PropsWithChildren) => {
  return (
    <SafeAreaView edges={['right', 'bottom', 'left']}>{children}</SafeAreaView>
  );
};

export default BaseLayout;
