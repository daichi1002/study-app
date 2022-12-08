export const formatDate = (date: string | null): string => {
  if (date == null) {
    return "";
  }
  const formatDate = new Date(date);
  return formatDate.toLocaleDateString();
};
