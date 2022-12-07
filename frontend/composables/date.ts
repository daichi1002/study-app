export const formatDate = (date: string) => {
  const formatDate = new Date(date);
  return formatDate.toLocaleDateString();
};
