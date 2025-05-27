interface Profile {
	nickname: string;
	handle: string;
	name: string;
	email: string;
	profession: string;
	salt: string;
}

interface Social {
	handle: string;
	url: string;
	name: string;
	description: string;
}

interface Experience {
	uuid: string;
	company: string;
	employment_type: string;
	location_type: string;
	position: string;
	start_date: string;
	end_date: string;
	summary: string;
	country: string;
	city: string;
	skills: string;
}

export interface User {
	profile: Profile;
	social: Social[];
	experience: Experience[];
}
