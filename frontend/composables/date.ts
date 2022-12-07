export const formatDate = (date: string | null) => {
  if (date == null) {
    return;
  }
  const formatDate = new Date(date);
  return formatDate.toLocaleDateString();
};
