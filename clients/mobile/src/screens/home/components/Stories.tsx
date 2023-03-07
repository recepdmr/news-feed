import React from 'react';
import {
  Image,
  ScrollView,
  StyleSheet,
  Text,
  TouchableOpacity,
  View,
} from 'react-native';
import useSWR from 'swr';

const fetcher = (...args: [RequestInfo, RequestInit | undefined]) =>
  fetch(...args)
    .then(res => res.json())
    .then(res => res.data);

type StoryId = string;

interface Story {
  profileImageUri: string;
  storyId: StoryId;
}
const apiBaseURL = 'https://fefed6ce-564d-4982-8370-386ec98694c0.mock.pstmn.io';

const Stories: React.FC = () => {
  const {data, isLoading} = useSWR(apiBaseURL + '/v1/stories', fetcher);

  const stories: Story[] = data;
  if (isLoading) {
    return <Text>Loading</Text>;
  }

  return (
    <ScrollView contentContainerStyle={styles.scrollViewContainer} horizontal>
      {stories?.map(item => (
        <StoryItem
          key={item.storyId}
          storyId={item.storyId}
          profileImageUri={item.profileImageUri}
        />
      ))}
    </ScrollView>
  );
};

const StoryItem: React.FC<Story> = ({storyId, profileImageUri}: Story) => (
  <View style={styles.story}>
    <TouchableOpacity onPress={() => openStory(storyId)}>
      <Image
        style={styles.storyImage}
        source={{
          uri: profileImageUri,
        }}
      />
    </TouchableOpacity>
  </View>
);

const openStory = (storyId: StoryId) => {
  console.log('opened story + ' + storyId);
};

const styles = StyleSheet.create({
  story: {
    flex: 1,
    width: 75,
    height: 'auto',
    borderRadius: 50,
  },
  scrollViewContainer: {
    margin: 20,
  },

  storyImage: {
    width: 50,
    height: 50,
    borderRadius: 50,
  },
});

export default Stories;
